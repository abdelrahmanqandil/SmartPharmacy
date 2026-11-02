package shared

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAuthRejectsMissingCredentials(t *testing.T) {
	handler := RequireAuth("secret", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	rr := httptest.NewRecorder()
	handler(rr, httptest.NewRequest(http.MethodPost, "/api/dispense", nil))

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusUnauthorized)
	}
}

func TestRequireAuthAcceptsBearerToken(t *testing.T) {
	handler := RequireAuth("secret", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, "/api/dispense", nil)
	req.Header.Set("Authorization", "Bearer secret")
	rr := httptest.NewRecorder()
	handler(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusNoContent)
	}
}
