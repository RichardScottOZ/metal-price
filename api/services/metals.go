package services

import (
	"context"
	"fmt"

	"github.com/chutified/metal-price/metal/protos/metal"
)

func GetPrice(mc metal.MetalClient, materialP string) (float64, error) {

	material, ok := metal.Materials_value[materialP]
	if !ok {
		return 0, fmt.Errorf("material %v not found", materialP)
	}
	fmt.Println(material)

	// create request
	request := &metal.MetalRequest{Metal: metal.Materials(material)}

	// call the service
	response, err := mc.GetPrice(context.Background(), request)
	if err != nil {
		return 0, err
	}

	// success
	return response.GetPrice(), nil
}
