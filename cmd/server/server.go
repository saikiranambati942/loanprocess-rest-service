package main

import (
	"loanprocess-rest-service/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/getbalance", handlers.GetBalance)
	http.HandleFunc("/payment", handlers.Payment)
	http.HandleFunc("/loaninitiate", handlers.LoanInitiate)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("Shutting down the application")
		os.Exit(1)
	}
}
