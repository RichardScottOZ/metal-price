package services

import "fmt"

var units = map[string]float64{
	"oz":  1,
	"lb":  16,
	"g":   0.0352739619,
	"dkg": 0.352739619,
	"kg":  35.2739619,
	"t":   32000,
}
var signs = map[string]string{
	"ounce":    "oz",
	"pound":    "lb",
	"gram":     "g",
	"decagram": "dkg",
	"kilogram": "kg",
	"ton":      "t",
}

// GetWeightRate returns the rate between two weight units.
func GetWeightRate(base, dest string) (float64, error) {

	// shorten
	if bb, ok := signs[base]; ok {
		base = bb
	}
	if dd, ok := signs[dest]; ok {
		dest = dd
	}

	// validation
	d, ok := units[dest]
	if !ok {
		return -1, fmt.Errorf("destination unit %s not found", dest)
	}
	b, ok := units[base]
	if !ok {
		return -1, fmt.Errorf("base unit %s not found", base)
	}

	// success
	rate := d / b
	return rate, nil
}
