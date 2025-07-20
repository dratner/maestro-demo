package handlers

import (
	"log"
	"net/http"
)

// Health handles the health check endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check endpoint accessed")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
