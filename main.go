package main

import (
	"log"
	"net/http"
)

// healthHandler handles the /health endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received health check request from %s", r.RemoteAddr)

	// Set content type header
	w.Header().Set("Content-Type", "text/plain")
	
	// Write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	// Configure the HTTP server
	port := ":8080"
	
	// Register the health check endpoint
	http.HandleFunc("/health", healthHandler)
	
	// Log server startup
	log.Printf("Starting server on port %s", port)
	log.Printf("Health check endpoint available at http://localhost%s/health", port)
	
	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
