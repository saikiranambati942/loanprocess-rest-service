package handlers

import (
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
}

func TestLoanInitiateHandlerNegative(t *testing.T) {

	t.Run("validating date format", func(t *testing.T) {
		x := `{
			"loanamount": 5000,
			"interest":5,
			"startdate":"2020/02/02"
			
		  }`
		r, _ := http.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		assert.Equal(t, w.Code, http.StatusBadRequest)
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
		assert.Equal(t, 400, w.Code)
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
