package services

import (
	"context"
	"fmt"

	"github.com/chutified/metal-price/metal/service/protos/metal"
)

// Metal handles the metal price services.
type Metal struct {
	client metal.MetalClient
}

// NewMetal is a constructor for the Metal service.
func NewMetal(mc metal.MetalClient) *Metal {
	return &Metal{
		client: mc,
	}
}

// GetPrice call the service and returns the price of the metal.
func (m *Metal) GetPrice(materialP string) (float64, error) {

	if mn, ok := PeriodicSymbols[materialP]; ok {
		materialP = mn
	}

	material, ok := metal.Materials_value[materialP]
	if !ok {
		return 0, fmt.Errorf("material %v not found", materialP)
	}

	// create request
	request := &metal.MetalRequest{Metal: metal.Materials(material)}

	// call the service
	response, err := m.client.GetPrice(context.Background(), request)
	if err != nil {
		return 0, err
	}

	// success
	return response.GetPrice(), nil
}
