package services

import (
	"context"

	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/gin-gonic/gin"
)

// GetRate provides the exchange rate of the base and destination.
func GetRate(c *gin.Context, baseP, destP string) (float32, error) {
	cc := c.Value("currency_client").(currency.CurrencyClient)

	base := currency.Currencies_value[baseP]
	dest := currency.Currencies_value[destP]

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
