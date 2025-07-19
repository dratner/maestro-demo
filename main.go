package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received health check request from %s", r.RemoteAddr)

	// Set content type header
	w.Header().Set("Content-Type", "text/plain")
	
	// Write response
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func main() {
	// Create a new serve mux
	mux := http.NewServeMux()
	
	// Register the health check endpoint
	mux.HandleFunc("/health", healthHandler)
	
	// Create server with reasonable timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)
	
	// Start the server
	go func() {
		log.Printf("Server listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()
	
	// Channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	
	// Blocking select waiting for either a server error or a shutdown signal
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)
		
	case sig := <-shutdown:
		log.Printf("Starting shutdown, received signal: %v", sig)
		
		// Create context with timeout for shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		
		// Attempt graceful shutdown
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop server gracefully: %v", err)
			if err := server.Close(); err != nil {
				log.Printf("Could not force close server: %v", err)
			}
		}
	}
}
