package node

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"smart-pharmacy/internal/config"
	"smart-pharmacy/internal/db"
	"smart-pharmacy/internal/inventory"
	"smart-pharmacy/internal/replica"
	"smart-pharmacy/internal/shared"
)

type edgeNode struct {
	cfg *config.Config
	db  *sql.DB
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

func requestLogger(role string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(sw, r)
		elapsed := time.Since(start)
		log.Printf("[%s] %s %s → %d  %s", role, r.Method, r.URL.RequestURI(), sw.status, elapsed.Round(time.Microsecond))
	})
}

func Start(cfg *config.Config) {
	conn := db.InitDB("edge", cfg.Database.Path)

	n := &edgeNode{
		cfg: cfg,
		db:  conn,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/status", n.handleStatus)
	mux.HandleFunc("/api/health", n.handleStatus)
	mux.HandleFunc("/api/dispense", shared.RequireAuth(cfg.Auth.Token, n.handleDispense))
	mux.HandleFunc("/api/medications", shared.RequireAuth(cfg.Auth.Token, n.handleMedications))
	mux.HandleFunc("/api/medications/register", shared.RequireAuth(cfg.Auth.Token, n.handleRegisterMedication))
	mux.HandleFunc("/api/inventory", shared.RequireAuth(cfg.Auth.Token, n.handleInventory))
	mux.HandleFunc("/api/transactions", shared.RequireAuth(cfg.Auth.Token, n.handleTransactions))
	mux.HandleFunc("/api/stock/add", shared.RequireAuth(cfg.Auth.Token, n.handleAddStock))
	mux.HandleFunc("/api/patients", shared.RequireAuth(cfg.Auth.Token, n.handlePatients))
	mux.HandleFunc("/api/patients/register", shared.RequireAuth(cfg.Auth.Token, n.handleRegisterPatient))
	mux.HandleFunc("/api/sync/export", shared.RequireAuth(cfg.Auth.Token, replica.ExportHandler(n.db)))
	mux.HandleFunc("/api/sync/import", shared.RequireAuth(cfg.Auth.Token, replica.ImportHandler(n.db)))

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("[EDGE] node=%s branch=%s listening on %s", cfg.Node.ID, cfg.Node.Branch, addr)

	log.Fatal(http.ListenAndServe(addr, shared.CORS(requestLogger("edge", mux))))
}

func (n *edgeNode) handleStatus(w http.ResponseWriter, _ *http.Request) {
	status := map[string]any{
		"node_id":      n.cfg.Node.ID,
		"role":         n.cfg.Node.Role,
		"branch":       n.cfg.Node.Branch,
		"capabilities": []string{"dispense", "inventory_read", "medication_register", "stock_add", "patient_create", "patient_update"},
		"timestamp":    time.Now().UTC().Format(time.RFC3339),
	}
	for k, v := range shared.ReplicationStatus(n.db, n.cfg.Database.Path) {
		status[k] = v
	}
	shared.Respond(w, status)
}

func (n *edgeNode) handleInventory(w http.ResponseWriter, _ *http.Request) {
	items, err := inventory.NewInventoryLedger(n.db).ListInventoryWithBranch(n.cfg.Node.Branch)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	result := []map[string]any{}
	for _, item := range items {
		result = append(result, map[string]any{
			"medication_id":  item.MedicationID,
			"name":           item.Name,
			"unit":           item.Unit,
			"min_stock":      item.MinStock,
			"quantity":       item.Balance,
			"balance":        item.Balance,
			"global_balance": item.GlobalBalance,
			"low_stock":      item.Balance < item.MinStock,
		})
	}
	shared.Respond(w, result)
}

func (n *edgeNode) handleMedications(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		n.handleRegisterMedication(w, r)
		return
	}
	if r.Method != http.MethodGet {
		shared.RespondErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	rows, err := n.db.Query(`SELECT id, name, IFNULL(unit,'tablet'), IFNULL(min_stock,0) FROM medications`)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer rows.Close()

	out := []map[string]any{}
	for rows.Next() {
		var id, name, unit string
		var minStock int
		if err := rows.Scan(&id, &name, &unit, &minStock); err != nil {
			shared.RespondErr(w, 500, err.Error())
			return
		}
		out = append(out, map[string]any{
			"medication_id": id,
			"name":          name,
			"unit":          unit,
			"min_stock":     minStock,
		})
	}
	shared.Respond(w, out)
}

func (n *edgeNode) handleRegisterMedication(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MedicationID string `json:"medication_id"`
		Name         string `json:"name"`
		Unit         string `json:"unit"`
		MinStock     int    `json:"min_stock"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.RespondErr(w, 400, "invalid JSON")
		return
	}
	if err := upsertMedication(n.db, req.MedicationID, req.Name, req.Unit, req.MinStock); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	shared.Respond(w, map[string]string{"status": "ok"})
}

func (n *edgeNode) handlePatients(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		n.handleRegisterPatient(w, r)
		return
	}
	if r.Method != http.MethodGet {
		shared.RespondErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	rows, err := n.db.Query(`SELECT id, name, IFNULL(date_of_birth,''), IFNULL(contact,''), IFNULL(branch_id,''), IFNULL(last_updated,'') FROM patients`)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer rows.Close()

	out := []map[string]any{}
	for rows.Next() {
		var id, name, dob, contact, branchID, lastUpdated string
		if err := rows.Scan(&id, &name, &dob, &contact, &branchID, &lastUpdated); err != nil {
			shared.RespondErr(w, 500, err.Error())
			return
		}
		out = append(out, map[string]any{
			"patient_id":    id,
			"name":          name,
			"date_of_birth": dob,
			"contact":       contact,
			"branch_id":     branchID,
			"last_updated":  lastUpdated,
		})
	}
	shared.Respond(w, out)
}

func (n *edgeNode) handleTransactions(w http.ResponseWriter, _ *http.Request) {
	rows, err := n.db.Query(`
		SELECT id, medication_id, IFNULL(patient_id,''), IFNULL(pharmacist_id,''),
		       quantity, type, IFNULL(branch_id,''), IFNULL(node_id,''),
		       IFNULL(status,'COMPLETED'), IFNULL(timestamp,'')
		FROM transactions
		ORDER BY rowid DESC
		LIMIT 200
	`)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer rows.Close()

	result := []map[string]any{}
	for rows.Next() {
		var id, medID, patID, pharmID, typ, branchID, nodeID, status, ts string
		var qty int
		rows.Scan(&id, &medID, &patID, &pharmID, &qty, &typ, &branchID, &nodeID, &status, &ts)
		result = append(result, map[string]any{
			"tx_id":         id,
			"medication_id": medID,
			"patient_id":    patID,
			"pharmacist_id": pharmID,
			"quantity":      qty,
			"type":          typ,
			"branch_id":     branchID,
			"node_id":       nodeID,
			"status":        status,
			"timestamp":     ts,
		})
	}
	shared.Respond(w, result)
}

func (n *edgeNode) handleAddStock(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MedicationID  string `json:"medication_id"`
		Quantity      int    `json:"quantity"`
		AddedBy       string `json:"added_by"`
		TransactionID string `json:"transaction_id"`
		LedgerEntryID string `json:"ledger_entry_id"`
		RecordedAt    string `json:"recorded_at"`
		NodeID        string `json:"node_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.RespondErr(w, 400, "invalid JSON")
		return
	}
	if req.MedicationID == "" || req.Quantity <= 0 {
		shared.RespondErr(w, 400, "medication_id and positive quantity are required")
		return
	}

	txID := req.TransactionID
	if txID == "" {
		txID = shared.NewID()
	}
	recordedAt := req.RecordedAt
	if recordedAt == "" {
		recordedAt = time.Now().UTC().Format("2006-01-02 15:04:05")
	}
	nodeID := req.NodeID
	if nodeID == "" {
		nodeID = n.cfg.Node.ID
	}

	tx, err := n.db.Begin()
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer tx.Rollback()

	if already, err := transactionExists(tx, txID); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	} else if already {
		if err := tx.Commit(); err != nil {
			shared.RespondErr(w, 500, err.Error())
			return
		}
		shared.Respond(w, map[string]string{"status": "ok", "tx_id": txID})
		return
	}

	if _, err := tx.Exec(`
		INSERT INTO transactions (id, medication_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp)
		VALUES (?, ?, ?, ?, 'STOCK_IN', ?, ?, 'COMPLETED', ?)
	`, txID, req.MedicationID, req.AddedBy, req.Quantity, n.cfg.Node.Branch, nodeID, recordedAt); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if err := inventory.RecordStockInWithID(tx, req.LedgerEntryID, req.MedicationID, req.Quantity, txID, nodeID, req.AddedBy, recordedAt); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	shared.Respond(w, map[string]string{"status": "ok", "tx_id": txID})
}

func (n *edgeNode) handleRegisterPatient(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PatientID   string `json:"patient_id"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
		Contact     string `json:"contact"`
		BranchID    string `json:"branch_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.RespondErr(w, 400, "invalid JSON")
		return
	}
	id := req.PatientID
	if id == "" {
		id = req.ID
	}
	branchID := req.BranchID
	if branchID == "" {
		branchID = n.cfg.Node.Branch
	}
	if err := upsertPatient(n.db, id, req.Name, req.DateOfBirth, req.Contact, branchID); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	shared.Respond(w, map[string]string{"status": "ok"})
}

type dispenseRequest struct {
	MedicationID string `json:"medication_id"`
	Quantity     int    `json:"quantity"`
	PatientID    string `json:"patient_id"`
	PharmacistID string `json:"pharmacist_id"`
}

func (n *edgeNode) handleDispense(w http.ResponseWriter, r *http.Request) {
	var req dispenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.RespondErr(w, 400, "invalid JSON")
		return
	}
	if req.MedicationID == "" || req.Quantity <= 0 {
		shared.RespondErr(w, 400, "medication_id and positive quantity are required")
		return
	}

	balance, err := inventory.NewInventoryLedger(n.db).ComputeCurrentInventory(req.MedicationID)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	if balance < req.Quantity {
		shared.RespondErr(w, 409, fmt.Sprintf("insufficient stock (have %d, need %d)", balance, req.Quantity))
		return
	}

	txID := shared.NewID()

	tx, err := n.db.Begin()
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO transactions (id, medication_id, patient_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp)
		VALUES (?, ?, ?, ?, ?, 'DISPENSE', ?, ?, 'COMPLETED', ?)
	`, txID, req.MedicationID, req.PatientID, req.PharmacistID, req.Quantity, n.cfg.Node.Branch, n.cfg.Node.ID, time.Now().UTC().Format("2006-01-02 15:04:05"))
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	if _, err := inventory.RecordDispenseIfAvailable(tx, req.MedicationID, req.Quantity, txID, n.cfg.Node.Branch, n.cfg.Node.ID, req.PharmacistID); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	log.Printf("[EDGE] DISPENSE tx=%s med=%-12s qty=%d patient=%s pharmacist=%s",
		txID, req.MedicationID, req.Quantity, req.PatientID, req.PharmacistID)

	shared.Respond(w, map[string]any{
		"tx_id":  txID,
		"status": "COMPLETED",
	})
}

func upsertPatient(db *sql.DB, id, name, dob, contact, branchID string) error {
	if id == "" || name == "" {
		return fmt.Errorf("patient_id and name are required")
	}

	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	var existing string
	err := db.QueryRow(`SELECT id FROM patients WHERE id = ?`, id).Scan(&existing)
	if err == sql.ErrNoRows {
		_, err = db.Exec(
			`INSERT INTO patients (id, name, date_of_birth, contact, branch_id, last_updated) VALUES (?, ?, ?, ?, ?, ?)`,
			id, name, dob, contact, branchID, now,
		)
		return err
	}
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`UPDATE patients SET name = ?, date_of_birth = ?, contact = ?, branch_id = ?, last_updated = ? WHERE id = ?`,
		name, dob, contact, branchID, now, id,
	)
	return err
}

func upsertMedication(db *sql.DB, id, name, unit string, minStock int) error {
	if id == "" || name == "" {
		return fmt.Errorf("medication_id and name are required")
	}
	if unit == "" {
		unit = "tablet"
	}

	var existing string
	err := db.QueryRow(`SELECT id FROM medications WHERE id = ?`, id).Scan(&existing)
	if err == sql.ErrNoRows {
		_, err = db.Exec(
			`INSERT INTO medications (id, name, unit, min_stock) VALUES (?, ?, ?, ?)`,
			id, name, unit, minStock,
		)
		return err
	}
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`UPDATE medications SET name = ?, unit = ?, min_stock = ? WHERE id = ?`,
		name, unit, minStock, id,
	)
	return err
}

func transactionExists(q interface {
	QueryRow(query string, args ...any) *sql.Row
}, id string) (bool, error) {
	var existing string
	err := q.QueryRow(`SELECT id FROM transactions WHERE id = ?`, id).Scan(&existing)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
