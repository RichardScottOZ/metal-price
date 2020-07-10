package data

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestNewRates(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)

	tests := []struct {
		name      string
		source    string
		expErrMsg string
	}{
		{
			name:      "ok",
			source:    "https://api.exchangeratesapi.io/latest",
			expErrMsg: "",
		},
		{
			name:      "invalid source",
			source:    "https://www.google.com",
			expErrMsg: "could not update exchange rates",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			r, err := NewRates(l, test.source)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.NotEqual(t1, r.log, nil)
				assert.NotEqual(t1, r.rates["CZK"], 0)
			}
		})
	}
}

func TestGetRare(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	r, err := NewRates(l, "https://api.exchangeratesapi.io/latest")
	if err != nil {
		t.Fatalf("unable to create new currency data service: %v", err)
	}

	tests := []struct {
		name      string
		base      string
		dest      string
		expErrMsg string
	}{
		{
			name:      "ok",
			base:      "EUR",
			dest:      "CZK",
			expErrMsg: "",
		},
		{
			name:      "invalid currency",
			base:      "EUR",
			dest:      "invalid",
			expErrMsg: "rate not found for currency",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			_, err := r.GetRate(test.base, test.dest)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
