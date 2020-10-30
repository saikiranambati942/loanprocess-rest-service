package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type payment struct {
	Repayment float64 `json:"amount"`
	PaidDate  string  `json:"date"`
}

// Payment function is a handler to handle the loan repayments
func Payment(w http.ResponseWriter, r *http.Request) {
	var p payment
	// Unmarshalling the json request data
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "input data format is not correct"}
		json.NewEncoder(w).Encode(er)
		return
	}
	// condition to check if the payment is less than or equal to zero
	if p.Repayment <= 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "amount should be greater than zero"}
		json.NewEncoder(w).Encode(er)
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	pd, err := Date(p.PaidDate) //payment date
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "date format should be YYYY-MM-DD"}
		json.NewEncoder(w).Encode(er)
		return
	}
	// condition to check if the requested add payment date is only after loan start date
	if pd.Before(lsd) || lsd.IsZero() {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "no loan record on this date"}
		json.NewEncoder(w).Encode(er)
		return
	}
	m := message{Message: "payment successful"}
	// condition to check whether there is any payment on the same date before.
	// If yes, instead of overwriting the previous payment we are adding the payments.
	if v, ok := datamap[pd]; ok {
		datamap[pd] = v + p.Repayment
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(m)
		return
	}
	// adding the payment date and amount to the map
	datamap[pd] = p.Repayment
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
