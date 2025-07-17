package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/username/app/pkg/handler"
)

func main() {
	h := handler.NewHandler()
	
	// Register routes
	http.HandleFunc("/", h.HelloWorld)
	
	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
