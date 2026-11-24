package health

import (
	"encoding/json"
	"net/http"

	"healthcheckservice/monitor"
)

// Handler returns an HTTP handler that writes the current health statuses as JSON.
func Handler(m *monitor.Monitor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statuses := m.GetStatuses()
		w.Header().Set("Content-Type", "application/json")
// noticed this could be clearer
		if err := json.NewEncoder(w).Encode(statuses); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}