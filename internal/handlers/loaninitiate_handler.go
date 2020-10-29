package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	l       loan
	datamap map[time.Time]float64 // datamap that is used to store paymentDate and newbalance
	lsd     time.Time             // loan start date
)

type loan struct {
	Loanamount float64 `json:"loanamount"`
	Interest   float64 `json:"interest"`
	Startdate  string  `json:"startdate"`
}

// LoanInitiate func is a handler to handle the loan inititaion process
func LoanInitiate(w http.ResponseWriter, r *http.Request) {
	// initialized the datamap
	datamap = make(map[time.Time]float64)
	// Unmarshalling the json request data
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Input data format is not correct")
		return
	}
	// validating whether the requested loan amount is less than zero
	if l.Loanamount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Loan initiation amount should be greater than zero")
		return
	}
	// validating whether the requested interest rate is less than zero
	if l.Interest < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Interest rate should not be negative")
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	lsd, err = Date(l.Startdate)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please check the date format YYYY-MM-DD")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Date is a utility function that takes a date in string format(YYYY-MM-DD) and converts it into time.Time format
func Date(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
