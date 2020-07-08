package handlers

import (
	"fmt"
	"strings"

	"github.com/chutified/metal-price/server/services"
	"github.com/gin-gonic/gin"
)

// GetPrice handles request of the price of the metal.
func (h *Handler) GetPrice(c *gin.Context) {

	// PARAMETERS
	// metal
	metal := c.Param("metal")
	metal = strings.ToLower(metal)
	// currency
	curr := c.Param("currency")
	curr = strings.ToLower(curr)
	// unit
	unit := c.Param("unit")
	unit = strings.ToLower(unit)

	// SERVICE CALLS
	// currency
	currRate, err := h.cs.GetRate("usd", curr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("unable to call currency service: %v", err),
		})
		return
	}
	// metal
	price, err := h.ms.GetPrice(metal) // ounces
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("call metal service: %v", err),
		})
		return
	}
	price *= float64(currRate)

	// SERVE

	// with undefined unit
	if unit == "" {
		c.JSON(200, gin.H{
			"metal":    metal,
			"price":    price,
			"currency": curr,
			"unit":     "oz",
		})
		return
	}

	// get unit's rate
	unitRate, err := services.GetWeightRate("oz", unit)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("call weight unit converter: %v", err),
		})
		return
	}
	price *= unitRate

	// success
	c.JSON(200, gin.H{
		"metal":    metal,
		"price":    price,
		"currency": curr,
		"unit":     unit,
	})
}
