package server

import (
	"context"
	"fmt"
	"log"

	"github.com/chutified/metal-price/metal/data"
	"github.com/chutified/metal-price/metal/protos/metal"
)

type Metal struct {
	log    *log.Logger
	prices *data.Prices
}

func NewMetal(l *log.Logger, pr *data.Prices) *Metal {
	return &Metal{
		log:    l,
		prices: pr,
	}
}

func (m *Metal) GetPrice(ctx context.Context, req *metal.MetalRequest) (*metal.MetalResponse, error) {

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
