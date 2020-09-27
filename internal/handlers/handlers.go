package handlers

import "net/http"

// Routes ...
func Routes() {
	http.HandleFunc("/getbalance", GetBalance)
	http.HandleFunc("/payment", Payment)
	http.HandleFunc("/loaninitiate", LoanInitiate)
}
