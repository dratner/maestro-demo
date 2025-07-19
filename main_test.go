package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	
	// Call the handler
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)
	
	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
	// Check Content-Type header
	expectedContentType := "text/plain"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}
	
	// Check response body
	expectedBody := "OK"
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}
	if got := strings.TrimSpace(string(body)); got != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expectedBody)
	}
}

func TestHealthHandlerIntegration(t *testing.T) {
	// Create a test server
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	server := httptest.NewServer(mux)
	defer server.Close()
	
	// Make a request to the test server
	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer resp.Body.Close()
	
	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	
	// Check Content-Type header
	expectedContentType := "text/plain"
	if contentType := resp.Header.Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}
	
	// Check response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}
	
	expectedBody := "OK"
	if got := strings.TrimSpace(string(body)); got != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expectedBody)
	}
}
