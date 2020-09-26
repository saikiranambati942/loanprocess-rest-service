package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"loanprocess-rest-service/utils"
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
	pd, err := utils.Date(p.PaidDate) //payment date
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please check the date format YYYY-MM-DD")
		return
	}
	numofDays := pd.Sub(lsd).Hours() / 24 // number of days elapsed between loan start date and particular payment date
	interestPerday := (l.Interest * b.newbalance) / (100 * 365)
	interestaccrued := interestPerday * float64(numofDays)

	//Assumed interest is calculated on daily basis
	//Checking if the same day has more than one payment. If yes, interest shouldn't be calculated again.
	if p.PaidDate == prevPaydate {
		interestaccrued = 0
	}
	prevPaydate = p.PaidDate
	b.newbalance = b.newbalance + interestaccrued - p.Repayment
	if b.newbalance < 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "paid extra amount than balance, you have credit balance of: %v ", -b.newbalance)
		return
	}
	b.date = pd
	m[pd] = b // ading the balance details of a particular day to map
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "payment done successfully")

}
