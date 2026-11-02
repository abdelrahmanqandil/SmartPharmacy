package inventory

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func TestInventoryLedgerComputesBalanceFromEntries(t *testing.T) {
	db := openLedgerTestDB(t)
	ledger := NewInventoryLedger(db)

	if err := ledger.RecordStockIn("MED-1", 10, "tx-stock", "central", "ADM-1"); err != nil {
		t.Fatalf("stock in: %v", err)
	}
	if _, err := ledger.DispenseIfAvailable("MED-1", 3, "tx-dispense", "branch-1", "edge-1", "PH-1"); err != nil {
		t.Fatalf("dispense: %v", err)
	}

	balance, err := ledger.ComputeCurrentInventory("MED-1")
	if err != nil {
		t.Fatalf("compute balance: %v", err)
	}
	if balance != 7 {
		t.Fatalf("balance = %d, want 7", balance)
	}

	items, err := ledger.ListInventory()
	if err != nil {
		t.Fatalf("list inventory: %v", err)
	}
	if len(items) != 1 || items[0].Balance != 7 {
		t.Fatalf("items = %#v, want one item with balance 7", items)
	}
}

func TestInventoryLedgerRejectsInsufficientStockWithoutAppending(t *testing.T) {
	db := openLedgerTestDB(t)
	ledger := NewInventoryLedger(db)

	if err := ledger.RecordStockIn("MED-1", 2, "tx-stock", "central", "ADM-1"); err != nil {
		t.Fatalf("stock in: %v", err)
	}
	if _, err := ledger.DispenseIfAvailable("MED-1", 3, "tx-dispense", "branch-1", "edge-1", "PH-1"); err == nil {
		t.Fatal("dispense succeeded with insufficient stock")
	}

	balance, err := ledger.ComputeCurrentInventory("MED-1")
	if err != nil {
		t.Fatalf("compute balance: %v", err)
	}
	if balance != 2 {
		t.Fatalf("balance = %d, want 2", balance)
	}
}

func openLedgerTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	for _, stmt := range []string{
		`CREATE TABLE medications (
			id TEXT PRIMARY KEY,
			name TEXT,
			unit TEXT,
			min_stock INTEGER
		)`,
		`CREATE TABLE inventory_ledger (
			id TEXT PRIMARY KEY,
			medication_id TEXT,
			quantity_delta INTEGER,
			entry_type TEXT,
			transaction_id TEXT,
			branch_id TEXT,
			node_id TEXT,
			recorded_by TEXT,
			timestamp TEXT
		)`,
		`INSERT INTO medications (id, name, unit, min_stock) VALUES ('MED-1', 'Aspirin', 'tablet', 5)`,
	} {
		if _, err := db.Exec(stmt); err != nil {
			t.Fatalf("exec %q: %v", stmt, err)
		}
	}
	return db
}
