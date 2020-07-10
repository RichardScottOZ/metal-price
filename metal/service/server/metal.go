package server

import (
	"context"
	"fmt"
	"log"

	config "github.com/chutified/metal-price/metal/config"
	data "github.com/chutified/metal-price/metal/service/data"
	metal "github.com/chutified/metal-price/metal/service/protos/metal"
)

// Metal is a the service server.
type Metal struct {
	log    *log.Logger
	prices *data.Prices
	cfg    *config.Config
}

// NewMetal constructs a new server.
func NewMetal(l *log.Logger, cfg *config.Config) *Metal {
	return &Metal{
		log: l,
		cfg: cfg,
	}
}

// GetPrice handles thegRPC request.
func (m *Metal) GetPrice(ctx context.Context, req *metal.MetalRequest) (*metal.MetalResponse, error) {

	// data service
	var err error
	m.prices, err = data.NewPrices(m.log, m.cfg.Source)
	if err != nil {
		return nil, fmt.Errorf("could not construct metal price data service: %w", err)
	}

	// get material
	material := req.GetMetal().String()

	// logging
	m.log.Printf("Handle GetPrice, material: %s\n", material)

	// get price
	price, err := m.prices.GetPrice(material)
	if err != nil {
		return nil, fmt.Errorf("unable to get the price of the material: %v", err)
	}

	// success
	metalResp := &metal.MetalResponse{Price: price}
	return metalResp, nil
}
