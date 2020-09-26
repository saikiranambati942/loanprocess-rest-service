package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

var b balance

type balanceDate struct {
	Date string `json:"date"`
}

type balance struct {
	newbalance float64
	date       time.Time
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
	if v, ok := m[dt]; ok {
		bal := strconv.FormatFloat(v.newbalance, 'f', 6, 64)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, bal)
		return
	}
	var days float64 = math.MaxFloat64
	var b float64
	for date, bal := range m {
		if date.Before(dt) {
			d := dt.Sub(date).Hours() / 24
			if d < days {
				days = d
				b = bal.newbalance
			}
		}
	}
	interestPerday := (l.Interest * b) / (100 * 365)
	interestaccrued := interestPerday * days
	balance := b + interestaccrued
	balanceString := strconv.FormatFloat(balance, 'f', 6, 64)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, balanceString)
}
