package services

import "fmt"

var units = map[string]float64{
	"oz": 1,
	"kg": 0.0283495231,
	"lb": 0.0625,
}

// GetWeightRate returns the rate between two weight units.
func GetWeightRate(base, dest string) (float64, error) {

	// validation
	d, ok := units[dest]
	if !ok {
		return -1, fmt.Errorf("unit %s not found", dest)
	}
	b, ok := units[base]
	if !ok {
		return -1, fmt.Errorf("unit %s not found", base)
	}

	// success
	rate := d / b
	return rate, nil
}
