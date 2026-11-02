package inventory

import (
	"database/sql"
	"fmt"
	"time"

	"smart-pharmacy/internal/shared"
)

type Item struct {
	MedicationID  string
	Name          string
	Unit          string
	MinStock      int
	Balance       int
	GlobalBalance int
}

type Entry struct {
	ID            string
	MedicationID  string
	QuantityDelta int
	EntryType     string
	TransactionID string
	BranchID      string
	NodeID        string
	RecordedBy    string
	Timestamp     string
}

type InventoryLedger struct {
	db *sql.DB
}

func NewInventoryLedger(db *sql.DB) *InventoryLedger {
	return &InventoryLedger{db: db}
}

func (l *InventoryLedger) ComputeCurrentInventory(medicationID string) (int, error) {
	return computeCurrentInventory(l.db, medicationID)
}

func (l *InventoryLedger) ListInventory() ([]Item, error) {
	rows, err := l.db.Query(`
		SELECT m.id, m.name, IFNULL(m.unit,'tablet'), IFNULL(m.min_stock,0), IFNULL(b.balance,0)
		FROM medications m
		LEFT JOIN (
			SELECT medication_id, SUM(quantity_delta) AS balance
			FROM inventory_ledger
			GROUP BY medication_id
		) b ON m.id = b.medication_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []Item{}
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.MedicationID, &item.Name, &item.Unit, &item.MinStock, &item.Balance); err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, rows.Err()
}

func (l *InventoryLedger) ListInventoryWithBranch(branchID string) ([]Item, error) {
	rows, err := l.db.Query(`
		SELECT m.id, m.name, IFNULL(m.unit,'tablet'), IFNULL(m.min_stock,0),
		       IFNULL(loc.balance,0), IFNULL(glob.balance,0)
		FROM medications m
		LEFT JOIN (
			SELECT medication_id, SUM(quantity_delta) AS balance
			FROM inventory_ledger
			WHERE branch_id = ?
			GROUP BY medication_id
		) loc ON m.id = loc.medication_id
		LEFT JOIN (
			SELECT medication_id, SUM(quantity_delta) AS balance
			FROM inventory_ledger
			GROUP BY medication_id
		) glob ON m.id = glob.medication_id
	`, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []Item{}
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.MedicationID, &item.Name, &item.Unit, &item.MinStock, &item.Balance, &item.GlobalBalance); err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, rows.Err()
}

func (l *InventoryLedger) RecordStockIn(medicationID string, quantity int, transactionID, nodeID, recordedBy string) error {
	return RecordStockIn(l.db, medicationID, quantity, transactionID, nodeID, recordedBy)
}

func RecordStockIn(e execer, medicationID string, quantity int, transactionID, nodeID, recordedBy string) error {
	return RecordStockInWithID(e, "", medicationID, quantity, transactionID, nodeID, recordedBy, "")
}

func RecordStockInWithID(e execer, entryID, medicationID string, quantity int, transactionID, nodeID, recordedBy, timestamp string) error {
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}
	if entryID == "" {
		entryID = shared.NewID()
	}
	if timestamp == "" {
		timestamp = nowSQL()
	}
	entry := Entry{
		ID:            entryID,
		MedicationID:  medicationID,
		QuantityDelta: quantity,
		EntryType:     "STOCK_IN",
		TransactionID: transactionID,
		NodeID:        nodeID,
		RecordedBy:    recordedBy,
		Timestamp:     timestamp,
	}
	return appendEntry(e, entry)
}

func (l *InventoryLedger) DispenseIfAvailable(medicationID string, quantity int, transactionID, branchID, nodeID, recordedBy string) (int, error) {
	tx, err := l.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	newBalance, err := RecordDispenseIfAvailable(tx, medicationID, quantity, transactionID, branchID, nodeID, recordedBy)
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return newBalance, nil
}

func RecordDispenseIfAvailable(e ledgerStore, medicationID string, quantity int, transactionID, branchID, nodeID, recordedBy string) (int, error) {
	if quantity <= 0 {
		return 0, fmt.Errorf("quantity must be greater than zero")
	}

	balance, err := computeCurrentInventory(e, medicationID)
	if err != nil {
		return 0, err
	}
	if balance < quantity {
		return balance, fmt.Errorf("insufficient stock (have %d, need %d)", balance, quantity)
	}

	entry := Entry{
		ID:            shared.NewID(),
		MedicationID:  medicationID,
		QuantityDelta: -quantity,
		EntryType:     "DISPENSE",
		TransactionID: transactionID,
		BranchID:      branchID,
		NodeID:        nodeID,
		RecordedBy:    recordedBy,
		Timestamp:     nowSQL(),
	}
	if err := appendEntry(e, entry); err != nil {
		return 0, err
	}
	return balance - quantity, nil
}

type queryer interface {
	QueryRow(query string, args ...any) *sql.Row
}

type execer interface {
	Exec(query string, args ...any) (sql.Result, error)
}

type ledgerStore interface {
	queryer
	execer
}

func computeCurrentInventory(q queryer, medicationID string) (int, error) {
	var balance int
	err := q.QueryRow(
		`SELECT IFNULL(SUM(quantity_delta),0) FROM inventory_ledger WHERE medication_id = ?`,
		medicationID,
	).Scan(&balance)
	return balance, err
}

func appendEntry(e execer, entry Entry) error {
	_, err := e.Exec(`
		INSERT INTO inventory_ledger
			(id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, entry.ID, entry.MedicationID, entry.QuantityDelta, entry.EntryType, entry.TransactionID, entry.BranchID, entry.NodeID, entry.RecordedBy, entry.Timestamp)
	return err
}

func nowSQL() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
