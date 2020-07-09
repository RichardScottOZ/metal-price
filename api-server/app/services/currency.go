package services

import (
	"context"
	"fmt"

	currency "github.com/chutified/metal-price/currency/service/protos/currency"
)

// Currency handles the currency services.
type Currency struct {
	client currency.CurrencyClient
}

// NewCurrency is a constructor for the Currency service.
func NewCurrency(cc currency.CurrencyClient) *Currency {
	return &Currency{
		client: cc,
	}
}

// GetRate call the service and returns the rate of two currencies.
func (c *Currency) GetRate(baseP, destP string) (float32, error) {

	base, ok := currency.Currencies_value[baseP]
	if !ok {
		return 0, fmt.Errorf("currency %s not found", baseP)
	}
	dest, ok := currency.Currencies_value[destP]
	if !ok {
		return 0, fmt.Errorf("currency %s not found", destP)
	}

	// request of the currency service
	request := &currency.RateRequest{
		Base:        currency.Currencies(base),
		Destination: currency.Currencies(dest),
	}

	// call currency service
	response, err := c.client.GetRate(context.Background(), request)
	if err != nil {
		return -1, err
	}

	// success
	return response.GetRate(), nil
}
