package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/thedevsaddam/gojsonq"
)

// Prices is a data service of metal price.
type Prices struct {
	log    *log.Logger
	prices map[string]float64
	source string
}

// NewPrices construct a new price data service.
func NewPrices(l *log.Logger, src string) (*Prices, error) {
	p := &Prices{
		log:    l,
		prices: map[string]float64{},
		source: src,
	}

	// update prices
	err := p.getPrices()
	if err != nil {
		return nil, fmt.Errorf("could not update metal prices: %w", err)
	}

	return p, nil
}

// GetPrice provides the price of the metal.
func (p *Prices) GetPrice(m string) (float64, error) {

	// get price
	price, ok := p.prices[m]
	if !ok {
		return 0, fmt.Errorf("material %s not found", m)
	}
	return price, nil
}

// getPrices udpates the map of prices in price data service
func (p *Prices) getPrices() error {

	body, err := pricesAPI(p.source)
	if err != nil {
		return fmt.Errorf("metal api error: %w", err)
	}

	// query material prices
	resp := gojsonq.New().FromString(string(body)).Get()
	m, ok := resp.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid source: %w", err)
	}

	for material, attributes := range m {
		material = strings.ToLower(material)

		// get body
		attrBody, ok := attributes.(map[string]interface{})
		if !ok {
			continue
		}

		// get price
		price, ok := attrBody["price"].(string)
		if !ok {
			continue
		}

		// filter
		price = strings.ReplaceAll(price, "$", "")
		price = strings.ReplaceAll(price, ",", "")

		p.prices[material], _ = strconv.ParseFloat(price, 64)
	}

	return nil
}

// pricesAPI returns the response body and the error
// from requesting a metal prices API.
func pricesAPI(api string) ([]byte, error) {

	// requesting
	resp, err := http.Get(api)
	if err != nil {
		return nil, fmt.Errorf("could not request the metal api: %w", err)
	}

	// check status code
	if resp.StatusCode != 200 {
		return nil, errors.New("expected status code 200")
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read metal api response body: %w", err)
	}
	defer resp.Body.Close()

	// success
	return body, nil
}
