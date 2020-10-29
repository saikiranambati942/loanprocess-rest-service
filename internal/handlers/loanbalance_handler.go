package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type balanceDate struct {
	Date string `json:"date"`
}

type balance struct {
	Balance string `json:"balance"`
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
	params := mux.Vars(r)
	bd.Date = params["date"]
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
	interestaccrued := 0.0
	// pd is the payment date. Ranging over the date slice
	for i, pd := range ds {
		var days float64
		if i == 0 {
			days = pd.Sub(lsd).Hours() / 24
		} else {
			days = pd.Sub(ds[i-1]).Hours() / 24
		}

		interestPerday := (l.Interest * b) / (100 * 365)
		interestaccrued = interestaccrued + (interestPerday * days)
		//new balance amount
		b = b - datamap[pd]
	}
	b = b + interestaccrued
	//converting the balance of float64 format t string format
	balanceString := strconv.FormatFloat(b, 'f', 6, 64)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	bal := balance{
		Balance: balanceString}
	json.NewEncoder(w).Encode(bal)

}
