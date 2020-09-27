package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	l   loan
	m   map[time.Time]float64 // map that is used to store paymentDate and newbalance
	lsd time.Time             // loan start date
)

type loan struct {
	Loanamount float64 `json:"loanamount"`
	Interest   float64 `json:"interest"`
	Startdate  string  `json:"startdate"`
}

// LoanInitiate func is a handler to handle the loan inititaion process
func LoanInitiate(w http.ResponseWriter, r *http.Request) {
	m = make(map[time.Time]float64)
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Input data format is not correct")
		return
	}
	if l.Loanamount <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Loan initiation amount should be greater than zero")
		return
	}
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
	fmt.Fprintf(w, "loan initiated successfully")
}

// Date is a utility function that takes a date in string format(YYYY-MM-DD) and converts it into time.Time format
func Date(date string) (time.Time, error) {
	d := strings.Split(date, "-")
	t := time.Time{} // zeroth value of time is nil struct
	year, err := strconv.Atoi(d[0])
	if err != nil {
		return t, err
	}
	month, err := strconv.Atoi(d[1])
	if err != nil {
		return t, err
	}
	day, err := strconv.Atoi(d[2])
	if err != nil {
		return t, err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}
