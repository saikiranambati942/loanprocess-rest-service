package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoanPaymentHandler(t *testing.T) {
	var m message
	var e errorResponse
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
		json.Unmarshal(w.Body.Bytes(), &m)
		assert.Equal(t, "payment successful", m.Message)
	})

	t.Run("testcase of loan payment zero", func(t *testing.T) {
		y := `{
			"amount": 0,
			"date":"2020-02-20"

		  }`
		r, _ := http.NewRequest("POST", "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "amount should be greater than zero", e.Error)
	})

	t.Run("testcase of loan payment before loan sanctioned date", func(t *testing.T) {
		y := `{
			"amount": 1000,
			"date":"2020-02-02"
   
		  }`
		r, _ := http.NewRequest("POST", "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "no loan record on this date", e.Error)

	})
	t.Run("validating date format", func(t *testing.T) {
		x := `{
			"amount": 1000,
			"date":"2020/02/02"
			
		  }`
		r, _ := http.NewRequest(http.MethodPost, "/payment", strings.NewReader(x))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		json.Unmarshal(w.Body.Bytes(), &e)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "date format should be YYYY-MM-DD", e.Error)
	})

}
