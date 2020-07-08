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
	log   *log.Logger
	rates map[string]float64
}

// NewRates returns an new empty data service.
func NewRates(l *log.Logger) (*Rates, error) {

	// contruct
	r := &Rates{
		log:   l,
		rates: map[string]float64{},
	}

	// update rates
	err := r.getRates("https://api.exchangeratesapi.io/latest")
	if err != nil {
		return nil, fmt.Errorf("could not update exchange rates: %w", err)
	}

	return r, nil
}

// GetRate returns rate of the two currencies.
func (r *Rates) GetRate(base, dest string) (float64, error) {

	// validation
	d, ok := r.rates[dest]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency: %s", dest)
	}
	b, ok := r.rates[base]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency: %s", base)
	}

	// calculate
	rate := d / b
	return rate, nil
}

// getRates updates exchange rates for the data service Rates.
func (r *Rates) getRates(api string) error {

	body, err := currencyAPI(api)
	if err != nil {
		return fmt.Errorf("currency api error: %w", err)
	}

	// query base and rates
	base := gojsonq.New().FromString(string(body)).Find("base").(string)
	rates := gojsonq.New().FromString(string(body)).From("rates").Get().(map[string]interface{})

	// set rates
	r.rates[base] = 1.0
	for c, rat := range rates {
		r.rates[c] = rat.(float64)
	}

	// success
	return nil
}

// currencyAPI returns the response body and the error
// from requesting an currency API.
func currencyAPI(api string) ([]byte, error) {

	// requesting
	resp, err := http.Get(api)
	if err != nil {
		return nil, fmt.Errorf("could not request the currency api: %w", err)
	}

	// check status code
	if resp.StatusCode != 200 {
		return nil, errors.New("expected status code 200")
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read currency api response body: %w", err)
	}
	defer resp.Body.Close()

	// success
	return body, nil
}
