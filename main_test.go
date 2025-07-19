package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)

	// Call the handler directly, passing in the request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expectedBody := "OK"
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expectedBody)
	}

	// Check the content type header
	expectedContentType := "text/plain"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}
}
