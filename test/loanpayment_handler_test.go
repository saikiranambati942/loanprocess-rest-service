package test

import (
	"loanprocess-rest-service/internal/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	handlers.Routes()
}

func TestLoanPaymentHandler(t *testing.T) {
	t.Run("positive testcase of statuscode check", func(t *testing.T) {
		x := `{
			"loanamount": 5000,
			"interest":5,
			"startdate":"2020-02-03"
			
		  }`
		r1 := httptest.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w1 := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w1, r1)
		if w1.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w1.Code)
		}
		y := `{
					 "amount": 2000,
					 "date":"2020-02-20"
			
				   }`
		r := httptest.NewRequest(http.MethodPost, "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w.Code)
		}
	})

	t.Run("testcase of loan payment zero", func(t *testing.T) {
		y := `{
			"amount": 0,
			"date":"2020-02-20"
   
		  }`
		r := httptest.NewRequest(http.MethodPost, "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		assert.Equal(t, "Please enter an amount greater than zero", w.Body.String())
	})

	t.Run("testcase of loan payment before loan sanctioned date", func(t *testing.T) {
		y := `{
			"amount": 1000,
			"date":"2020-02-02"
   
		  }`
		r := httptest.NewRequest(http.MethodPost, "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		assert.Equal(t, "There is no loan record on this date", w.Body.String())
		assert.Equal(t, 400, w.Code)

	})

}

func TestDate(t *testing.T) {

	t.Run("validating positive testcase of Date utility function", func(t *testing.T) {
		d, _ := handlers.Date("2020-02-02")
		expected := "2020-02-02 00:00:00 +0000 UTC"
		assert.Equal(t, expected, d.String())

	})

	t.Run("validating negative testcase of Date utility function", func(t *testing.T) {
		_, err := handlers.Date("2020-0@-02")
		assert.Error(t, err)
	})

}
