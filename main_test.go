package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GET request returns OK",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "POST request returns method not allowed",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req := httptest.NewRequest(tt.method, "/health", nil)
			w := httptest.NewRecorder()

			// Call handler
			healthHandler(w, req)

			// Check status code
			if got := w.Code; got != tt.expectedStatus {
				t.Errorf("healthHandler() status = %v, want %v", got, tt.expectedStatus)
			}

			// Check response body
			body, _ := io.ReadAll(w.Body)
			if got := string(body); got != tt.expectedBody {
				t.Errorf("healthHandler() body = %v, want %v", got, tt.expectedBody)
			}

			// For successful GET requests, verify Content-Type header
			if tt.method == http.MethodGet && tt.expectedStatus == http.StatusOK {
				if got := w.Header().Get("Content-Type"); got != "text/plain" {
					t.Errorf("healthHandler() Content-Type = %v, want text/plain", got)
				}
			}
		})
	}
}
