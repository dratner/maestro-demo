package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the health check endpoint
	http.HandleFunc("/health", handlers.Health)

	// Start the server
	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
