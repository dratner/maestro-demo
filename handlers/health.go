package handlers

import (
	"log"
	"net/http"
)

// Health handles the health check endpoint
func Health(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received health check request from %s", r.RemoteAddr)

	// Set content type header
	w.Header().Set("Content-Type", "text/plain")
	
	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
