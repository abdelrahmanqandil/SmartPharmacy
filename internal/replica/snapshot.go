package replica

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"smart-pharmacy/internal/shared"
)

type Snapshot struct {
	Medications  []MedicationRow  `json:"medications"`
	Ledger       []LedgerRow      `json:"ledger"`
	Transactions []TransactionRow `json:"transactions"`
	Patients     []PatientRow     `json:"patients"`
}

type MedicationRow struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	MinStock int    `json:"min_stock"`
}

type LedgerRow struct {
	ID            string `json:"id"`
	MedicationID  string `json:"medication_id"`
	QuantityDelta int    `json:"quantity_delta"`
	EntryType     string `json:"entry_type"`
	TransactionID string `json:"transaction_id"`
	BranchID      string `json:"branch_id"`
	NodeID        string `json:"node_id"`
	RecordedBy    string `json:"recorded_by"`
	Timestamp     string `json:"timestamp"`
}

type TransactionRow struct {
	ID           string `json:"id"`
	MedicationID string `json:"medication_id"`
	PatientID    string `json:"patient_id"`
	PharmacistID string `json:"pharmacist_id"`
	Quantity     int    `json:"quantity"`
	Type         string `json:"type"`
	BranchID     string `json:"branch_id"`
	NodeID       string `json:"node_id"`
	Status       string `json:"status"`
	Timestamp    string `json:"timestamp"`
}

type PatientRow struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Contact     string `json:"contact"`
	BranchID    string `json:"branch_id"`
	LastUpdated string `json:"last_updated"`
}

func ExportHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		snapshot, err := Export(db)
		if err != nil {
			shared.RespondErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		shared.Respond(w, snapshot)
	}
}

func ImportHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var snapshot Snapshot
		if err := json.NewDecoder(r.Body).Decode(&snapshot); err != nil {
			shared.RespondErr(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		counts, err := Import(db, snapshot)
		if err != nil {
			shared.RespondErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		shared.Respond(w, counts)
	}
}

func Export(db *sql.DB) (Snapshot, error) {
	var snapshot Snapshot
	var err error

	snapshot.Medications, err = exportMedications(db)
	if err != nil {
		return snapshot, err
	}
	snapshot.Ledger, err = exportLedger(db)
	if err != nil {
		return snapshot, err
	}
	snapshot.Transactions, err = exportTransactions(db)
	if err != nil {
		return snapshot, err
	}
	snapshot.Patients, err = exportPatients(db)
	if err != nil {
		return snapshot, err
	}

	return snapshot, nil
}

func Import(db *sql.DB, snapshot Snapshot) (map[string]int, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	counts := map[string]int{
		"medications":  0,
		"ledger":       0,
		"transactions": 0,
		"patients":     0,
	}

	for _, row := range snapshot.Medications {
		if err := upsertMedication(tx, row); err != nil {
			return nil, err
		}
		counts["medications"]++
	}
	for _, row := range snapshot.Transactions {
		inserted, err := insertTransactionIfMissing(tx, row)
		if err != nil {
			return nil, err
		}
		if inserted {
			counts["transactions"]++
		}
	}
	for _, row := range snapshot.Ledger {
		inserted, err := insertLedgerIfMissing(tx, row)
		if err != nil {
			return nil, err
		}
		if inserted {
			counts["ledger"]++
		}
	}
	for _, row := range snapshot.Patients {
		if err := upsertPatient(tx, row); err != nil {
			return nil, err
		}
		counts["patients"]++
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return counts, nil
}

func exportMedications(db *sql.DB) ([]MedicationRow, error) {
	rows, err := db.Query(`SELECT id, name, IFNULL(unit,'tablet'), IFNULL(min_stock,0) FROM medications`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []MedicationRow{}
	for rows.Next() {
		var row MedicationRow
		if err := rows.Scan(&row.ID, &row.Name, &row.Unit, &row.MinStock); err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

func exportLedger(db *sql.DB) ([]LedgerRow, error) {
	rows, err := db.Query(`
		SELECT id, medication_id, quantity_delta, entry_type, IFNULL(transaction_id,''),
		       IFNULL(branch_id,''), IFNULL(node_id,''), IFNULL(recorded_by,''), IFNULL(timestamp,'')
		FROM inventory_ledger
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []LedgerRow{}
	for rows.Next() {
		var row LedgerRow
		if err := rows.Scan(&row.ID, &row.MedicationID, &row.QuantityDelta, &row.EntryType, &row.TransactionID, &row.BranchID, &row.NodeID, &row.RecordedBy, &row.Timestamp); err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

func exportTransactions(db *sql.DB) ([]TransactionRow, error) {
	rows, err := db.Query(`
		SELECT id, medication_id, IFNULL(patient_id,''), IFNULL(pharmacist_id,''), quantity,
		       type, IFNULL(branch_id,''), IFNULL(node_id,''), IFNULL(status,'COMPLETED'), IFNULL(timestamp,'')
		FROM transactions
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []TransactionRow{}
	for rows.Next() {
		var row TransactionRow
		if err := rows.Scan(&row.ID, &row.MedicationID, &row.PatientID, &row.PharmacistID, &row.Quantity, &row.Type, &row.BranchID, &row.NodeID, &row.Status, &row.Timestamp); err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

func exportPatients(db *sql.DB) ([]PatientRow, error) {
	rows, err := db.Query(`
		SELECT id, name, IFNULL(date_of_birth,''), IFNULL(contact,''), IFNULL(branch_id,''), IFNULL(last_updated,'')
		FROM patients
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := []PatientRow{}
	for rows.Next() {
		var row PatientRow
		if err := rows.Scan(&row.ID, &row.Name, &row.DateOfBirth, &row.Contact, &row.BranchID, &row.LastUpdated); err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

func upsertMedication(tx *sql.Tx, row MedicationRow) error {
	var existing string
	err := tx.QueryRow(`SELECT id FROM medications WHERE id = ?`, row.ID).Scan(&existing)
	if err == sql.ErrNoRows {
		_, err = tx.Exec(`INSERT INTO medications (id, name, unit, min_stock) VALUES (?, ?, ?, ?)`, row.ID, row.Name, row.Unit, row.MinStock)
		return err
	}
	if err != nil {
		return err
	}
	_, err = tx.Exec(`UPDATE medications SET name = ?, unit = ?, min_stock = ? WHERE id = ?`, row.Name, row.Unit, row.MinStock, row.ID)
	return err
}

func upsertPatient(tx *sql.Tx, row PatientRow) error {
	var existing string
	err := tx.QueryRow(`SELECT id FROM patients WHERE id = ?`, row.ID).Scan(&existing)
	if err == sql.ErrNoRows {
		_, err = tx.Exec(
			`INSERT INTO patients (id, name, date_of_birth, contact, branch_id, last_updated) VALUES (?, ?, ?, ?, ?, ?)`,
			row.ID, row.Name, row.DateOfBirth, row.Contact, row.BranchID, row.LastUpdated,
		)
		return err
	}
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`UPDATE patients SET name = ?, date_of_birth = ?, contact = ?, branch_id = ?, last_updated = ? WHERE id = ?`,
		row.Name, row.DateOfBirth, row.Contact, row.BranchID, row.LastUpdated, row.ID,
	)
	return err
}

func insertTransactionIfMissing(tx *sql.Tx, row TransactionRow) (bool, error) {
	var existing string
	err := tx.QueryRow(`SELECT id FROM transactions WHERE id = ?`, row.ID).Scan(&existing)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == nil {
		return false, nil
	}
	_, err = tx.Exec(
		`INSERT INTO transactions (id, medication_id, patient_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		row.ID, row.MedicationID, row.PatientID, row.PharmacistID, row.Quantity, row.Type, row.BranchID, row.NodeID, row.Status, row.Timestamp,
	)
	return err == nil, err
}

func insertLedgerIfMissing(tx *sql.Tx, row LedgerRow) (bool, error) {
	var existing string
	err := tx.QueryRow(`SELECT id FROM inventory_ledger WHERE id = ?`, row.ID).Scan(&existing)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == nil {
		return false, nil
	}
	_, err = tx.Exec(
		`INSERT INTO inventory_ledger (id, medication_id, quantity_delta, entry_type, transaction_id, branch_id, node_id, recorded_by, timestamp)
		 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		row.ID, row.MedicationID, row.QuantityDelta, row.EntryType, row.TransactionID, row.BranchID, row.NodeID, row.RecordedBy, row.Timestamp,
	)
	return err == nil, err
}
