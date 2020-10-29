package handlers

import (
	"github.com/gorilla/mux"
)

// Routes function routes requests to a specific handler based the requestendpoint
func Routes(r *mux.Router) {
	r.HandleFunc("/loaninitiate", LoanInitiate).Methods("POST")
	r.HandleFunc("/payment", Payment).Methods("POST")
	r.HandleFunc("/getbalance/{date}", GetBalance).Methods("GET")
}
