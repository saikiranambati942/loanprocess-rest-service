package handlers

import (
	"encoding/json"
	"fmt"
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
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Input data format is not correct")
		return
	}
	// condition to check if the payment is less than or equal to zero
	if p.Repayment <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please enter an amount greater than zero")
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	pd, err := Date(p.PaidDate) //payment date
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please check the date format YYYY-MM-DD")
		return
	}
	// condition to check if the payment date requested to add is before loan initiated date
	if pd.Before(lsd) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "There is no loan record on this date")
		return
	}
	// condition to check whether there is any payment on the same date before.
	// If yes, instead of overwriting the previous payment we are adding the payments.
	if v, ok := datamap[pd]; ok {
		datamap[pd] = v + p.Repayment
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "payment done successfully")
		return
	}
	// adding the payment date and amount to the map
	datamap[pd] = p.Repayment
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "payment done successfully")
}
