package server

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

type Server struct {
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
	conn := db.InitDB(cfg.Node.Role, cfg.Database.Path)

	s := &Server{
		cfg: cfg,
		db:  conn,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/status", s.handleStatus)
	mux.HandleFunc("/api/medications", shared.RequireAuth(cfg.Auth.Token, s.handleMedications))
	mux.HandleFunc("/api/patients", shared.RequireAuth(cfg.Auth.Token, s.handlePatients))
	mux.HandleFunc("/api/inventory", shared.RequireAuth(cfg.Auth.Token, s.handleInventory))
	mux.HandleFunc("/api/transactions", shared.RequireAuth(cfg.Auth.Token, s.handleTransactions))
	mux.HandleFunc("/api/medications/register", shared.RequireAuth(cfg.Auth.Token, s.handleRegisterMedication))
	mux.HandleFunc("/api/patients/register", shared.RequireAuth(cfg.Auth.Token, s.handleRegisterPatient))
	mux.HandleFunc("/api/stock/add", shared.RequireAuth(cfg.Auth.Token, s.handleAddStock))
	mux.HandleFunc("/api/dispense", shared.RequireAuth(cfg.Auth.Token, s.handleDispense))
	mux.HandleFunc("/api/sync/export", shared.RequireAuth(cfg.Auth.Token, replica.ExportHandler(s.db)))
	mux.HandleFunc("/api/sync/import", shared.RequireAuth(cfg.Auth.Token, replica.ImportHandler(s.db)))

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("[%s] node=%s listening on %s", cfg.Node.Role, cfg.Node.ID, addr)

	log.Fatal(http.ListenAndServe(addr, shared.CORS(requestLogger(cfg.Node.Role, mux))))
}

func (s *Server) handleStatus(w http.ResponseWriter, _ *http.Request) {
	status := map[string]any{
		"node_id": s.cfg.Node.ID,
		"role":    s.cfg.Node.Role,
		"time":    time.Now().UTC(),
	}
	for k, v := range shared.ReplicationStatus(s.db, s.cfg.Database.Path) {
		status[k] = v
	}
	shared.Respond(w, status)
}

func (s *Server) handleMedications(w http.ResponseWriter, _ *http.Request) {
	rows, err := s.db.Query(`SELECT id, name, IFNULL(unit,'tablet'), IFNULL(min_stock,0) FROM medications`)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer rows.Close()

	out := []map[string]any{}
	for rows.Next() {
		var id, name, unit string
		var minStock int
		rows.Scan(&id, &name, &unit, &minStock)
		out = append(out, map[string]any{
			"medication_id": id,
			"name":          name,
			"unit":          unit,
			"min_stock":     minStock,
		})
	}
	shared.Respond(w, out)
}

func (s *Server) handlePatients(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		s.handleRegisterPatient(w, r)
		return
	}
	if r.Method != http.MethodGet {
		shared.RespondErr(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	rows, err := s.db.Query(`SELECT id, name, IFNULL(date_of_birth,''), IFNULL(contact,''), IFNULL(branch_id,''), IFNULL(last_updated,'') FROM patients`)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer rows.Close()

	out := []map[string]any{}
	for rows.Next() {
		var id, name, dob, contact, branchID, lastUpdated string
		rows.Scan(&id, &name, &dob, &contact, &branchID, &lastUpdated)
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

func (s *Server) handleInventory(w http.ResponseWriter, r *http.Request) {
	branch := r.URL.Query().Get("branch")

	out := []map[string]any{}
	if branch != "" {
		items, err := inventory.NewInventoryLedger(s.db).ListInventoryWithBranch(branch)
		if err != nil {
			shared.RespondErr(w, 500, err.Error())
			return
		}
		for _, item := range items {
			out = append(out, map[string]any{
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
	} else {
		items, err := inventory.NewInventoryLedger(s.db).ListInventory()
		if err != nil {
			shared.RespondErr(w, 500, err.Error())
			return
		}
		for _, item := range items {
			out = append(out, map[string]any{
				"medication_id": item.MedicationID,
				"name":          item.Name,
				"unit":          item.Unit,
				"min_stock":     item.MinStock,
				"quantity":      item.Balance,
				"balance":       item.Balance,
				"low_stock":     item.Balance < item.MinStock,
			})
		}
	}
	shared.Respond(w, out)
}

func (s *Server) handleTransactions(w http.ResponseWriter, _ *http.Request) {
	rows, err := s.db.Query(`
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

	out := []map[string]any{}
	for rows.Next() {
		var id, medID, patID, pharmID, typ, branchID, nodeID, status, ts string
		var qty int
		rows.Scan(&id, &medID, &patID, &pharmID, &qty, &typ, &branchID, &nodeID, &status, &ts)
		out = append(out, map[string]any{
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
	shared.Respond(w, out)
}

func (s *Server) handleRegisterMedication(w http.ResponseWriter, r *http.Request) {
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
	if req.Unit == "" {
		req.Unit = "tablet"
	}
	if err := upsertMedication(s.db, req.MedicationID, req.Name, req.Unit, req.MinStock); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	shared.Respond(w, map[string]string{"status": "ok"})
}

func (s *Server) handleRegisterPatient(w http.ResponseWriter, r *http.Request) {
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
	if err := upsertPatient(s.db, id, req.Name, req.DateOfBirth, req.Contact, req.BranchID); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	shared.Respond(w, map[string]string{"status": "ok"})
}

func (s *Server) handleAddStock(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MedicationID  string `json:"medication_id"`
		Quantity      int    `json:"quantity"`
		AddedBy       string `json:"added_by"`
		TransactionID string `json:"transaction_id"`
		LedgerEntryID string `json:"ledger_entry_id"`
		RecordedAt    string `json:"recorded_at"`
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
	tx, err := s.db.Begin()
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
		INSERT INTO transactions (id, medication_id, pharmacist_id, quantity, type, node_id, status, timestamp)
		VALUES (?, ?, ?, ?, 'STOCK_IN', ?, 'COMPLETED', ?)
	`, txID, req.MedicationID, req.AddedBy, req.Quantity, s.cfg.Node.ID, recordedAt); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if err := inventory.RecordStockInWithID(tx, req.LedgerEntryID, req.MedicationID, req.Quantity, txID, s.cfg.Node.ID, req.AddedBy, recordedAt); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	shared.Respond(w, map[string]string{"status": "ok", "tx_id": txID})
}

type dispenseRequest struct {
	MedicationID   string `json:"medication_id"`
	Quantity       int    `json:"quantity"`
	PatientID      string `json:"patient_id"`
	PharmacistID   string `json:"pharmacist_id"`
	BranchOverride string `json:"branch_override"`
}

func (s *Server) handleDispense(w http.ResponseWriter, r *http.Request) {
	var req dispenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		shared.RespondErr(w, 400, "invalid JSON")
		return
	}
	if req.MedicationID == "" || req.Quantity <= 0 {
		shared.RespondErr(w, 400, "medication_id and positive quantity are required")
		return
	}

	balance, err := inventory.NewInventoryLedger(s.db).ComputeCurrentInventory(req.MedicationID)
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if balance < req.Quantity {
		shared.RespondErr(w, 409, fmt.Sprintf("insufficient stock (have %d, need %d)", balance, req.Quantity))
		return
	}

	txID := shared.NewID()
	branchID := req.BranchOverride
	if branchID == "" {
		if s.cfg.Node.Branch != "" {
			branchID = s.cfg.Node.Branch
		} else {
			branchID = "central"
		}
	}

	tx, err := s.db.Begin()
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO transactions (id, medication_id, patient_id, pharmacist_id, quantity, type, branch_id, node_id, status, timestamp)
		VALUES (?, ?, ?, ?, ?, 'DISPENSE', ?, ?, 'COMPLETED', ?)
	`, txID, req.MedicationID, req.PatientID, req.PharmacistID, req.Quantity, branchID, s.cfg.Node.ID, time.Now().UTC().Format("2006-01-02 15:04:05"))
	if err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	if _, err := inventory.RecordDispenseIfAvailable(tx, req.MedicationID, req.Quantity, txID, branchID, s.cfg.Node.ID, req.PharmacistID); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}
	if err := tx.Commit(); err != nil {
		shared.RespondErr(w, 500, err.Error())
		return
	}

	log.Printf("[%s] DISPENSE tx=%s med=%-12s qty=%d branch=%s patient=%s",
		s.cfg.Node.Role, txID, req.MedicationID, req.Quantity, branchID, req.PatientID)

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
