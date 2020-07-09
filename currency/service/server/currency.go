package server

import (
	"context"
	"fmt"
	"log"

	"github.com/chutified/metal-price/currency/service/data"
	"github.com/chutified/metal-price/currency/service/protos/currency"
)

// Currency is a currency server.
type Currency struct {
	log   *log.Logger
	rates *data.Rates
}

// NewCurrency is a contructor for the Currency server.
func NewCurrency(l *log.Logger, r *data.Rates) *Currency {
	return &Currency{
		log:   l,
		rates: r,
	}
}

// GetRate returns a exchange rate of the request's base and destination currencies.
func (c *Currency) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {

	// get currencies
	base := req.GetBase().String()
	dest := req.GetDestination().String()

	// logging
	c.log.Printf("Handle GetRate, base: %s, destination: %s\n", base, dest)

	// get the rate
	rate, err := c.rates.GetRate(base, dest)
	if err != nil {
		return nil, fmt.Errorf("could not call GetRate currency service: %w", err)
	}

	// success
	rateResp := currency.RateResponse{Rate: float32(rate)}
	return &rateResp, nil
}
