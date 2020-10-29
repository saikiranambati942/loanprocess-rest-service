package handlers

import (
	"encoding/json"
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
		Router().ServeHTTP(w1, r1)
		if w1.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w1.Code)
		}
		y := `{
					 "amount": 1000,
					 "date":"2020-02-20"
			
				   }`
		r := httptest.NewRequest(http.MethodPost, "/payment", strings.NewReader(y))
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		if w.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w.Code)
		}
		r2 := httptest.NewRequest("GET", "/getbalance/2020-02-20", nil)
		w2 := httptest.NewRecorder()
		Router().ServeHTTP(w2, r2)
		if w2.Code != 200 {
			t.Fatalf("should receive a statuscode of %d but received %d", http.StatusOK, w2.Code)
		}
	})

	t.Run("validating the balance on a payment date", func(t *testing.T) {

		r := httptest.NewRequest("GET", "/getbalance/2020-02-20", nil)
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		expectedBalance := "4012.328767"
		var b balance
		json.Unmarshal(w.Body.Bytes(), &b)
		assert.Equal(t, expectedBalance, b.Balance)

	})
	t.Run("validating the balance on a non payment date", func(t *testing.T) {

		r := httptest.NewRequest("GET", "/getbalance/2020-02-22", nil)
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		expectedBalance := "4013.424658"
		var b balance
		json.Unmarshal(w.Body.Bytes(), &b)
		assert.Equal(t, expectedBalance, b.Balance)
	})

	t.Run("validating the balance before loan start date", func(t *testing.T) {

		r := httptest.NewRequest("GET", "/getbalance/2020-02-01", nil)
		w := httptest.NewRecorder()
		Router().ServeHTTP(w, r)
		expected := "There is no loan record on this date"
		assert.Equal(t, expected, w.Body.String())
		assert.Equal(t, 400, http.StatusBadRequest)
	})

}
