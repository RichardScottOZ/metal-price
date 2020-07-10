package data

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestNewPrices(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)

	tests := []struct {
		name      string
		source    string
		expErrMsg string
	}{
		{
			name:      "ok",
			source:    "https://www.moneymetals.com/api/spot-prices.json",
			expErrMsg: "",
		},
		{
			name:      "invalid source",
			source:    "https://www.google.com",
			expErrMsg: "could not update metal prices",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			p, err := NewPrices(l, test.source)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.NotEqual(t1, p.log, nil)
				assert.NotEqual(t1, p.prices["gold"], 0)
			}
		})
	}
}

func TestGetPrice(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	p, err := NewPrices(l, "https://www.moneymetals.com/api/spot-prices.json")
	if err != nil {
		t.Fatalf("unable to create new price data service: %v", err)
	}

	tests := []struct {
		name      string
		metal     string
		expErrMsg string
	}{
		{
			name:      "ok",
			metal:     "rhodium",
			expErrMsg: "",
		},
		{
			name:      "invalid metal",
			metal:     "invalid",
			expErrMsg: "material .* not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			_, err := p.GetPrice(test.metal)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
