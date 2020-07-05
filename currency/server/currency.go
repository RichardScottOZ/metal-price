package server

import (
	"context"
	"log"

	"github.com/chutified/metal-value/currency/protos/currency"
)

// Currency is a currency server.
type Currency struct {
	logger *log.Logger
}

// New is a contructor for the Currency server.
func New(l *log.Logger) *Currency {
	return &Currency{
		logger: l,
	}
}

// GetRate returns a exchange rate of the request's base and destination currencies.
func (c *Currency) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {
	strBase := currency.Currencies_name[int32(req.GetBase())]
	strDest := currency.Currencies_name[int32(req.GetDestination())]
	c.logger.Printf("Handler GetRate, base: %s, destination: %s\n", strBase, strDest)

	res := currency.RateResponse{
		Rate: 0.5,
	}

	return &res, nil
}
