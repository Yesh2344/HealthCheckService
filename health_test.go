package health

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"healthcheckservice/config"
	"healthcheckservice/monitor"
)

func TestHealthHandler(t *testing.T) {
	// Create two test servers: one healthy (200) and one unhealthy (500).
	healthySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer healthySrv.Close()

	unhealthySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer unhealthySrv.Close()

	// Prepare configuration pointing at the test servers.
	cfg := &config.Config{
		Services:      []string{healthySrv.URL, unhealthySrv.URL},
		CheckInterval: 100 * time.Millisecond,
	}
	mon := monitor.New(cfg)
	mon.Start()
	// Give the monitor a moment to perform its first round of checks.
	time.Sleep(200 * time.Millisecond)

	// Perform request against the health endpoint.
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	handler := Handler(mon)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	body := rr.Body.String()
	if !strings.Contains(body, healthySrv.URL) || !strings.Contains(body, "true") {
		t.Errorf("expected healthy service (%s) to be true in response: %s", healthySrv.URL, body)
	}
	if !strings.Contains(body, unhealthySrv.URL) || !strings.Contains(body, "false") {
		t.Errorf("expected unhealthy service (%s) to be false in response: %s", unhealthySrv.URL, body)
	}
	mon.Stop()
}