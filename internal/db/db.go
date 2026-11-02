package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

func InitDB(role, sqlitePath string) *sql.DB {
	dir := filepath.Dir(sqlitePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("[db] failed to create directory %s: %v", dir, err)
	}

	db, err := sql.Open("sqlite", sqlitePath)
	if err != nil {
		log.Fatalf("[db] failed to open SQLite at %s: %v", sqlitePath, err)
	}

	db.SetMaxOpenConns(1)
	if err := db.Ping(); err != nil {
		log.Fatalf("[db] SQLite ping failed: %v", err)
	}

	db.Exec(`PRAGMA journal_mode=WAL`)
	db.Exec(`PRAGMA foreign_keys=ON`)
	db.Exec(`PRAGMA busy_timeout=5000`)

	log.Printf("[db] open · role=%s · path=%s · mode=WAL", role, sqlitePath)

	var count int
	db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='medications'").Scan(&count)
	if count > 0 {
		log.Printf("[db] schema ok · skipping migration")
	} else {
		log.Printf("[db] schema missing · running migration...")
		initSchema(db)
	}

	waitForSchema(db)

	return db
}

func waitForSchema(db *sql.DB) {
	log.Printf("[db] verifying schema...")
	for {
		var count int
		if err := db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='medications'").Scan(&count); err == nil && count > 0 {
			log.Printf("[db] schema verified · ready")
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func initSchema(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS medications (
			id        TEXT PRIMARY KEY,
			name      TEXT,
			unit      TEXT,
			min_stock INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS inventory (
			medication_id TEXT PRIMARY KEY,
			quantity      INTEGER
		)`,
		`CREATE TABLE IF NOT EXISTS inventory_ledger (
			id             TEXT PRIMARY KEY,
			medication_id  TEXT,
			quantity_delta INTEGER,
			entry_type      TEXT,
			transaction_id  TEXT,
			branch_id       TEXT,
			node_id         TEXT,
			recorded_by     TEXT,
			timestamp       TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS transactions (
			id            TEXT PRIMARY KEY,
			medication_id TEXT,
			patient_id    TEXT,
			pharmacist_id TEXT,
			quantity      INTEGER,
			type          TEXT,
			branch_id     TEXT,
			node_id       TEXT,
			status        TEXT,
			timestamp     TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS patients (
			id            TEXT PRIMARY KEY,
			name          TEXT,
			date_of_birth TEXT,
			contact       TEXT,
			branch_id     TEXT,
			last_updated  TEXT
		)`,
	}
	for i, q := range tables {
		execDDL(db, q, i+1, len(tables))
		time.Sleep(300 * time.Millisecond)
	}
	migrateInventorySnapshot(db)
	log.Printf("[db] migration complete")
}

func migrateInventorySnapshot(db *sql.DB) {
	rows, err := db.Query(`SELECT medication_id, quantity FROM inventory`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var medicationID string
		var quantity int
		if err := rows.Scan(&medicationID, &quantity); err != nil {
			continue
		}
		if quantity == 0 {
			continue
		}

		var ledgerEntries int
		if err := db.QueryRow(`SELECT COUNT(*) FROM inventory_ledger WHERE medication_id = ?`, medicationID).Scan(&ledgerEntries); err != nil {
			continue
		}
		if ledgerEntries > 0 {
			continue
		}

		db.Exec(`
			INSERT INTO inventory_ledger
				(id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp)
			VALUES (?, ?, ?, 'LEGACY_SNAPSHOT', '', '', 'migration', 'migration', ?)
		`, "legacy-"+medicationID, medicationID, quantity, time.Now().UTC().Format("2006-01-02 15:04:05"))
	}
}

func execDDL(db *sql.DB, q string, n, total int) {
	for attempt := 1; attempt <= 10; attempt++ {
		_, err := db.Exec(q)
		if err == nil {
			log.Printf("[db] DDL %d/%d applied", n, total)
			return
		}
		msg := err.Error()
		if strings.Contains(msg, "duplicate column") ||
			strings.Contains(msg, "Duplicate column") ||
			strings.Contains(msg, "partial commit") {
			log.Printf("[db] DDL %d/%d already applied · skipping", n, total)
			return
		}
		log.Printf("[db] DDL %d/%d attempt %d failed: %v — retrying...", n, total, attempt, err)
		time.Sleep(500 * time.Millisecond)
	}
	log.Printf("[db] DDL %d/%d could not complete after 10 attempts — continuing", n, total)
}
