package handlers

import (
	"fmt"
	"math"
	"strings"

	services "github.com/chutommy/metal-price/api-server/app/services"
	"github.com/gin-gonic/gin"
)

// GetMetalMC handles request of the price of the metal.
// @summary Get a price in ounces (metal, currency).
// @description Get a price of the metal in a certain currency.
// @produce json
// @param metal path string true "the whole chemical name or an abbreviated version of a chemical element"
// @param currency path string true "the currency acronym"
// @success 200 {object} handlers.Response "ok"
// @failure 400 {object} handlers.HTTPError "service call"
// @router /i/{metal}/{currency} [get]
func (h *Handler) GetMetalMC(c *gin.Context) {

	// PARAMETERS
	// metal
	metal := c.Param("metal")
	metal = strings.ToLower(metal)
	if mn, ok := services.PeriodicSymbols[metal]; ok {
		metal = mn
	}
	// currency
	curr := c.Param("currency")
	curr = strings.ToUpper(curr)

	// SERVICE CALLS
	// currency
	currRate, err := h.cs.GetRate("USD", curr)
	if err != nil {
		c.JSON(400, &HTTPError{
			Message: fmt.Sprintf("call currency service: %v", err),
		})
		return
	}
	// metal
	price, err := h.ms.GetPrice(metal) // ounces
	if err != nil {
		c.JSON(400, &HTTPError{
			Message: fmt.Sprintf("call metal service: %v", err),
		})
		return
	}
	price *= float64(currRate)

	price = math.Round(price*100) / 100
	// success
	c.JSON(200, &Response{
		Metal:    metal,
		Price:    price,
		Currency: curr,
		Unit:     "oz",
	})

}
