package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thedevsaddam/gojsonq"
)

// Rates defines the exchange rates.
type Rates struct {
	log  *log.Logger
	rate map[string]float64
}

// NewRates returns an new empty data service.
func NewRates(l *log.Logger) (*Rates, error) {
	r := &Rates{
		log:  l,
		rate: map[string]float64{},
	}
	r.getRates()

	return r, nil
}

// GetRate returns rate of the two currencies.
func (r *Rates) GetRate(base, dest string) (float64, error) {

	// validation
	d, ok := r.rate[dest]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency: %s", dest)
	}
	b, ok := r.rate[base]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency: %s", base)
	}

	// success
	rate := d / b
	return rate, nil
}

// getRates updates exchange rates for the data service Rates.
func (r *Rates) getRates() error {

	// requesting
	resp, err := http.Get("https://api.exchangeratesapi.io/latest")
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("expected error code 200")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// query base and rates
	base := gojsonq.New().FromString(string(body)).Find("base").(string)
	rates := gojsonq.New().FromString(string(body)).From("rates").Get().(map[string]interface{})

	// set rates
	r.rate[base] = 1.0
	for c, rat := range rates {
		r.rate[c] = rat.(float64)
	}

	// success
	return nil
}
