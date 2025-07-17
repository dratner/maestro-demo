package handler

import (
	"fmt"
	"net/http"
)

// Handler holds dependencies for all handlers
type Handler struct{}

// NewHandler creates a new handler instance
func NewHandler() *Handler {
	return &Handler{}
}

// HelloWorld handles the root endpoint
func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
