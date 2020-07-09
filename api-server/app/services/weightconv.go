package services

import "fmt"

var units = map[string]float64{
	"oz": 1,
	"kg": 35.2739619,
	"lb": 16,
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
