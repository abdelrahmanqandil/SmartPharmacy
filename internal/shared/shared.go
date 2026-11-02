package shared

import (
	"crypto/subtle"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func Respond(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func RespondErr(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Pharmacy-Token")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireAuth(token string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if token == "" {
			RespondErr(w, http.StatusInternalServerError, "authentication token is not configured")
			return
		}
		if !validToken(r, token) {
			RespondErr(w, http.StatusUnauthorized, "authentication required")
			return
		}
		next(w, r)
	}
}

func validToken(r *http.Request, expected string) bool {
	actual := r.Header.Get("X-Pharmacy-Token")
	if actual == "" {
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, "Bearer ") {
			actual = strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		}
	}
	return subtle.ConstantTimeCompare([]byte(actual), []byte(expected)) == 1
}

func ReplicationStatus(db *sql.DB, sqlitePath string) map[string]any {
	state := "unknown"
	details := "Marmot sidecar status is not exposed to the application"

	if err := db.Ping(); err != nil {
		state = "unavailable"
		details = err.Error()
	} else if sqlitePath != "" {
		walPath := sqlitePath + "-wal"
		if info, err := os.Stat(walPath); err == nil {
			state = "wal_ready"
			details = fmt.Sprintf("local WAL present; last modified %s", info.ModTime().UTC().Format(time.RFC3339))
		} else if os.IsNotExist(err) {
			state = "wal_not_present"
			details = "local WAL file is not present yet"
		} else {
			details = err.Error()
		}
	}

	return map[string]any{
		"replication_state":   state,
		"replication_details": details,
	}
}

func NewID() string {
	return fmt.Sprintf("%d-%04x", time.Now().UnixNano(), rand.Intn(0xffff))
}
