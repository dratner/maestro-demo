package main

import (
	"log"
	"net/http"
)

// healthHandler handles the /health endpoint requests
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header
	w.Header().Set("Content-Type", "text/plain")
	
	// Write response with 200 OK status
	w.WriteHeader(http.StatusOK)
	
	// Write body
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	// Register the health check endpoint
	http.HandleFunc("/health", healthHandler)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
