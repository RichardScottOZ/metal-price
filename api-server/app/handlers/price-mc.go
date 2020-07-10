package handlers

import (
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetMetalMC handles request of the price of the metal.
func (h *Handler) GetMetalMC(c *gin.Context) {

	// PARAMETERS
	// metal
	metal := c.Param("metal")
	metal = strings.ToLower(metal)
	// currency
	curr := c.Param("currency")
	curr = strings.ToUpper(curr)

	// SERVICE CALLS
	// currency
	currRate, err := h.cs.GetRate("USD", curr)
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

	price = math.Round(price*100) / 100
	// success
	c.JSON(200, &response{
		Metal:    metal,
		Price:    price,
		Currency: curr,
		Unit:     "oz",
	})

}
