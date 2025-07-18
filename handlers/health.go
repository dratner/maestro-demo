package handlers

import (
	"net/http"
)

// HealthHandler responds to HTTP health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
