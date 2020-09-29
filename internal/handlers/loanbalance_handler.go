package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type balanceDate struct {
	Date string `json:"date"`
}

type dateSlice []time.Time

// Implemented Interface methods to sort the dataSilce
func (d dateSlice) Len() int {
	return len(d)
}
func (d dateSlice) Less(i, j int) bool {
	return d[i].Before(d[j])
}

func (d dateSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// GetBalance function is a handler to handle the requests to know the amount of totalbalance remaining on a specific date
func GetBalance(w http.ResponseWriter, r *http.Request) {
	var bd balanceDate
	// Unmarshalling the json request data
	err := json.NewDecoder(r.Body).Decode(&bd)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Input data format is not correct")
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	bldate, err := Date(bd.Date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please check the date format YYYY-MM-DD")
		return
	}
	// condition to check if the requested date to get balance is only after loan start date
	if bldate.Before(lsd) || lsd.IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "There is no loan record on this date")
		return
	}
	// checking whether the requested balance date is in the map of paid dates list.
	// If not, adding a entry in map with key as requested balance date and value as zero
	if _, ok := datamap[bldate]; !ok {
		datamap[bldate] = 0
	}
	ds := make(dateSlice, 0, len(datamap))
	// adding all the map keys which are dates to a slice
	for k := range datamap {
		ds = append(ds, k)
	}
	sort.Sort(ds) // sorting the slice of dates
	for i, paidDate := range ds {
		// reslicing the slice till the requested balance date
		if paidDate.Equal(bldate) {
			ds = ds[:i+1]
		}
	}
	// Assigning the initial balance to loan amount
	b := l.Loanamount
	// pd is the payment date. Ranging over the date slice
	for _, pd := range ds {
		// number of days between loan start date and payment date
		days := pd.Sub(lsd).Hours() / 24
		interestPerday := (l.Interest * b) / (100 * 365)
		interestaccrued := interestPerday * days
		//new balance amount
		b = b + interestaccrued - datamap[pd]
	}
	//converting the balance of float64 format t string format
	balanceString := strconv.FormatFloat(b, 'f', 6, 64)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Your Loan Balance as of %s is %s", bd.Date, balanceString)
}
