package handlers

import (
	"fmt"
	"math"
	"strings"

	services "github.com/chutified/metal-price/api-server/app/services"
	"github.com/gin-gonic/gin"
)

// GetMetalMCU handles request of the price of the metal.
// @summary Get a price (metal, currency, weight unit).
// @description Get a price of the metal in a certain currency and weight unit.
// @produce json
// @param metal path string true "the whole chemical name or an abbreviated version of a chemical element"
// @param currency path string true "the currency acronym"
// @param unit path string true "weight unit"
// @success 200 {object} handlers.Response "ok"
// @failure 400 {object} handlers.HTTPError "service call"
// @router /i/{metal}/{currency}/{unit} [get]
func (h *Handler) GetMetalMCU(c *gin.Context) {

	// PARAMETERS
	// metal
	metal := c.Param("metal")
	metal = strings.ToLower(metal)
	// currency
	curr := c.Param("currency")
	curr = strings.ToUpper(curr)
	// unit
	unit := c.Param("unit")
	unit = strings.ToLower(unit)

	// SERVICE CALLS
	// currency
	currRate, err := h.cs.GetRate("USD", curr)
	if err != nil {
		c.JSON(400, &HTTPError{
			Message: fmt.Sprintf("unable to call currency service: %v", err),
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

	// SERVE
	// get unit's rate
	unitRate, err := services.GetWeightRate("oz", unit)
	if err != nil {
		c.JSON(400, &HTTPError{
			Message: fmt.Sprintf("call weight unit converter: %v", err),
		})
		return
	}
	price *= unitRate

	price = math.Round(price*100) / 100
	// success
	c.JSON(200, &Response{
		Metal:    metal,
		Price:    price,
		Currency: curr,
		Unit:     unit,
	})
}
