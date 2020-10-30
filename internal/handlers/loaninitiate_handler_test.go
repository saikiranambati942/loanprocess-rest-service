package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/loaninitiate", LoanInitiate).Methods("POST")
	r.HandleFunc("/payment", Payment).Methods("POST")
	r.HandleFunc("/balance/{date}", GetBalance).Methods("GET")
	return r
}

func TestLoanInitiateHandlerPositive(t *testing.T) {
	var m message
	x := `{
		"loanamount": 5000,
		"interest":5,
		"startdate":"2020-02-02"
		
	  }`
	r, _ := http.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
	w := httptest.NewRecorder()
	Router().ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w.Code)
	}
	json.Unmarshal(w.Body.Bytes(), &m)
	assert.Equal(t, "loan initiation successful", m.Message)
}

func TestLoanInitiateHandlerNegative(t *testing.T) {
	var e errorResponse
	t.Run("validating date format", func(t *testing.T) {
		x := `{
			"loanamount": 5000,
			"interest":5,
			"startdate":"2020/02/02"
			
		  }`
		r, _ := http.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "date format should be YYYY-MM-DD", e.Error)
	})

	t.Run("validating input request format", func(t *testing.T) {
		//malformed input
		x := `{
			"loanamount": 5000,
			"interest":5             
			"startdate":"2020-02-02"
			
		  }`
		r, _ := http.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "input data format is not correct", e.Error)
	})

	t.Run("validating when loan amount less than or equal to zero", func(t *testing.T) {
		x := `{
			"loanamount": 0,
			"interest":5,             
			"startdate":"2020-02-02"
			
		  }`
		r, _ := http.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "loan initiation amount should be greater than zero", e.Error)
	})

	t.Run("validating when interest is less than zero", func(t *testing.T) {
		x := `{
			"loanamount": 5000,
			"interest": -5,             
			"startdate":"2020-02-02"
			
		  }`
		r, _ := http.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "interest rate cannot be negative", e.Error)
	})

}

func TestDate(t *testing.T) {
	t.Run("validating positive testcase of Date utility function", func(t *testing.T) {
		d, _ := Date("2020-02-02")
		expected := "2020-02-02 00:00:00 +0000 UTC"
		assert.Equal(t, expected, d.String())

	})

	t.Run("validating negative testcase of Date utility function", func(t *testing.T) {
		_, err := Date("2020-0@-02")
		assert.Error(t, err)
	})

}
