package services

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetWeightRate(t *testing.T) {

	tests := []struct {
		name   string
		base   string
		dest   string
		rate   float64
		errMsg string
	}{
		{
			name:   "ok",
			base:   "oz",
			dest:   "kg",
			rate:   35.27396,
			errMsg: "",
		},
		{
			name:   "ok full name",
			base:   "ounce",
			dest:   "kilogram",
			rate:   35.27396,
			errMsg: "",
		},
		{
			name:   "cross converting",
			base:   "g",
			dest:   "kg",
			rate:   1000,
			errMsg: "",
		},
		{
			name:   "invalid base",
			base:   "invalid",
			dest:   "oz",
			rate:   0,
			errMsg: "base unit .* not found",
		},
		{
			name:   "invalid destination",
			base:   "oz",
			dest:   "invalid",
			rate:   0,
			errMsg: "destination unit .* not found",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			r, err := GetWeightRate(test.base, test.dest)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.errMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				r := math.Round(r*100000) / 100000

				assert.Equal(t1, r, test.rate)
				assert.Equal(t1, "", test.errMsg)
			}
		})
	}
}
