package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoanPaymentHandler(t *testing.T) {
	t.Run("positive testcase of statuscode check", func(t *testing.T) {
		x := `{
			"loanamount": 5000,
			"interest":5,
			"startdate":"2020-02-03"
			
		  }`
		r1, _ := http.NewRequest("POST", "/loaninitiate", strings.NewReader(x))
		w1 := httptest.NewRecorder()
		Router().ServeHTTP(w1, r1)
		if w1.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w1.Code)
		}
		y := `{
					 "amount": 2000,
					 "date":"2020-02-20"
			
				   }`
		r, _ := http.NewRequest("POST", "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		if w.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w.Code)
		}
	})

	// t.Run("testcase of loan payment zero", func(t *testing.T) {
	// 	y := `{
	// 		"amount": 0,
	// 		"date":"2020-02-20"

	// 	  }`
	// 	r, _ := http.NewRequest("POST", "/payment", strings.NewReader(y))
	// 	w := httptest.NewRecorder()
	// 	Router().ServeHTTP(w, r)
	// 	assert.Equal(t, "Please enter an amount greater than zero", w.Body.String())
	// })

	t.Run("testcase of loan payment before loan sanctioned date", func(t *testing.T) {
		y := `{
			"amount": 1000,
			"date":"2020-02-02"
   
		  }`
		r, _ := http.NewRequest("POST", "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		assert.Equal(t, 400, w.Code)

	})

}
