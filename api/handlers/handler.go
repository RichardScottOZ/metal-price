package handlers

import (
	"log"

	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/chutified/metal-price/metal/protos/metal"
)

// Handler defines the api handler.
type Handler struct {
	log *log.Logger
	cc  currency.CurrencyClient
	mc  metal.MetalClient
}

// NewHandler is the constructor of the Handler.
func NewHandler(l *log.Logger, cc currency.CurrencyClient, mc metal.MetalClient) *Handler {
	return &Handler{
		log: l,
		cc:  cc,
		mc:  mc,
	}
}
