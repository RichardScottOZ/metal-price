package handlers

import (
	"fmt"
	"math"
	"strings"

	services "github.com/chutified/metal-price/api-server/app/services"
	"github.com/gin-gonic/gin"
)

// GetMetalM handles request of the price of the metal.
// @summary Get a price in USD pre an ounce (metal).
// @description Get a price of the metal.
// @produce json
// @param metal path string true "the whole chemical name or an abbreviated version of a chemical element"
// @success 200 {object} handlers.Response "ok"
// @failure 400 {object} handlers.HTTPError "service call"
// @router /i/{metal} [get]
func (h *Handler) GetMetalM(c *gin.Context) {

	// PARAMETERS
	// metal
	metal := c.Param("metal")
	metal = strings.ToLower(metal)
	if mn, ok := services.PeriodicSymbols[metal]; ok {
		metal = mn
	}

	// SERVICE CALLS
	// metal
	price, err := h.ms.GetPrice(metal) // ounces
	if err != nil {
		c.JSON(400, &HTTPError{
			Message: fmt.Sprintf("call metal service: %v", err),
		})
		return
	}

	price = math.Round(price*100) / 100
	// success
	c.JSON(200, &Response{
		Metal:    metal,
		Price:    price,
		Currency: "USD",
		Unit:     "oz",
	})

}
