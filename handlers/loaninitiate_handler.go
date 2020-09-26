package handlers

import (
	"encoding/json"
	"fmt"
	"loanprocess-rest-service/utils"
	"log"
	"net/http"
	"time"
)

type loan struct {
	Loanamount float64 `json:"loanamount"`
	Interest   float64 `json:"interest"`
	Startdate  string  `json:"startdate"`
}

// LoanInitiate func is a handler to handle the loan inititaion process
func LoanInitiate(w http.ResponseWriter, r *http.Request) {
	m = make(map[time.Time]balance)
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Input data format is not correct")
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	lsd, err = utils.Date(l.Startdate)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please check the date format YYYY-MM-DD")
		return
	}
	b.newbalance = l.Loanamount
	b.date = lsd
	m[lsd] = b
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "loan initiated successfully")
}
