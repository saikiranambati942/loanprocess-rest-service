package main

import (
	"loanprocess-rest-service/internal/handlers"
	"log"
	"net/http"
	"os"
)

// Starting point of the application
func main() {
	handlers.Routes()
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("Shutting down the application")
		os.Exit(1)
	}
}
