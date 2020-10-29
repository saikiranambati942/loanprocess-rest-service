package main

import (
	"loanprocess-rest-service/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// application starts here
func main() {
	r := mux.NewRouter()
	handlers.Routes(r)
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal("Shutting down the application")
		os.Exit(1)
	}
}
