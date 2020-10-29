package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var (
	l       loan
	datamap map[time.Time]float64 // datamap that is used to store paymentDate and newbalance
	lsd     time.Time             // loan start date
)

type message struct {
	Message string `json:"message"`
}
type errorResponse struct {
	Error string `json:"error"`
}
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
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "input data format is not correct"}
		json.NewEncoder(w).Encode(er)
		return
	}
	// validating whether the requested loan amount is less than zero
	if l.Loanamount <= 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "loan initiation amount should be greater than zero"}
		json.NewEncoder(w).Encode(er)
		return
	}
	// validating whether the requested interest rate is less than zero
	if l.Interest < 0 {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "interest rate cannot be negative"}
		json.NewEncoder(w).Encode(er)
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	lsd, err = Date(l.Startdate)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		er := errorResponse{
			Error: "date format should be YYYY-MM-DD"}
		json.NewEncoder(w).Encode(er)
		return
	}
	m := message{Message: "loan initiation successful"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

// Date is a utility function that takes a date in string format(YYYY-MM-DD) and converts it into time.Time format
func Date(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}
