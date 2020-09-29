package handlers

import "net/http"

// Routes function routes requests to a specific handler based the requestendpoint
func Routes() {
	http.HandleFunc("/getbalance", GetBalance)
	http.HandleFunc("/payment", Payment)
	http.HandleFunc("/loaninitiate", LoanInitiate)
}
