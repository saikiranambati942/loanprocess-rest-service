package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoanBalanceHandler(t *testing.T) {
	t.Run("positive testcase of statuscode", func(t *testing.T) {
		x := `{
			"loanamount": 5000,
			"interest":5,
			"startdate":"2020-02-02"
			
		  }`
		r1 := httptest.NewRequest(http.MethodPost, "/loaninitiate", strings.NewReader(x))
		w1 := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w1, r1)
		if w1.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w1.Code)
		}
		y := `{
					 "amount": 1000,
					 "date":"2020-02-20"
			
				   }`
		r := httptest.NewRequest(http.MethodPost, "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w.Code)
		}
		z := `{
			"date":"2020-02-20"
		  }`
		r2 := httptest.NewRequest(http.MethodPost, "/getbalance", strings.NewReader(z))
		w2 := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w2, r2)
		if w2.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w2.Code)
		}

	})

	t.Run("validating the balance on a payment date", func(t *testing.T) {
		x := `{
			"date":"2020-02-20"
		  }`
		r := httptest.NewRequest(http.MethodPost, "/getbalance", strings.NewReader(x))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		expectedBalance := "4012.328767"
		assert.Equal(t, expectedBalance, w.Body.String())

	})
	t.Run("validating the balance on a non payment date", func(t *testing.T) {
		x := `{
			"date":"2020-02-22"
		  }`
		r := httptest.NewRequest(http.MethodPost, "/getbalance", strings.NewReader(x))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		expectedBalance := "4023.321449"
		assert.Equal(t, expectedBalance, w.Body.String())
	})

	t.Run("validating the balance before loan start date", func(t *testing.T) {
		x := `{
			"date":"2020-02-01"
		  }`
		r := httptest.NewRequest(http.MethodPost, "/getbalance", strings.NewReader(x))
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)
		expected := "There is no loan record on this date"
		assert.Equal(t, expected, w.Body.String())
		assert.Equal(t, 400, http.StatusBadRequest)
	})

}
