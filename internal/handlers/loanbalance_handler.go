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
	err := json.NewDecoder(r.Body).Decode(&bd)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Input data format is not correct")
		return
	}
	// Assuming date in the request is in the format of YYYY-MM-DD
	dt, err := Date(bd.Date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please check the date format YYYY-MM-DD")
		return
	}
	if dt.Before(lsd) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "There is no loan record on this date")
		return
	}
	if _, ok := m[dt]; !ok {
		m[dt] = 0
	}
	ds := make(dateSlice, 0, len(m))
	for k := range m {
		ds = append(ds, k)
	}
	sort.Sort(ds)
	for k, v := range ds {
		if v.Equal(dt) {
			ds = ds[:k+1]
		}
	}
	b := l.Loanamount
	for _, pd := range ds {
		days := pd.Sub(lsd).Hours() / 24
		interestPerday := (l.Interest * b) / (100 * 365)
		interestaccrued := interestPerday * days
		b = b + interestaccrued - m[pd]
	}
	balanceString := strconv.FormatFloat(b, 'f', 6, 64)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, balanceString)
}
