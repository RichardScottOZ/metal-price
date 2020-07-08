package services

import (
	"context"
	"fmt"

	"github.com/chutified/metal-price/currency/protos/currency"
)

// GetRate provides the exchange rate of the base and destination.
func GetRate(cc currency.CurrencyClient, baseP, destP string) (float32, error) {

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
	response, err := cc.GetRate(context.Background(), request)
	if err != nil {
		return -1, err
	}

	// success
	return response.GetRate(), nil
}
