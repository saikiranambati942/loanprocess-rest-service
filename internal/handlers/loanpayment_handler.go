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
	if pd.Before(lsd) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "There is no loan record on this date")
		return
	}
	if v, ok := m[pd]; ok {
		m[pd] = v + p.Repayment
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "payment done successfully")
		return
	}
	m[pd] = p.Repayment // ading the balance details of a particular day to map
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "payment done successfully")
}
