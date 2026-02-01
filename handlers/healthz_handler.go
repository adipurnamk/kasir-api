package handlers

import (
    "net/http"
)

// HealthzHandler handles GET requests to the /healthz endpoint
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"OK", "message": "API Running"}`))
}