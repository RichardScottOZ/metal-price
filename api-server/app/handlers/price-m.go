package handlers

import (
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetMetalM handles request of the price of the metal.
func (h *Handler) GetMetalM(c *gin.Context) {

	// PARAMETERS
	// metal
	metal := c.Param("metal")
	metal = strings.ToLower(metal)

	// SERVICE CALLS
	// metal
	price, err := h.ms.GetPrice(metal) // ounces
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("call metal service: %v", err),
		})
		return
	}

	price = math.Round(price*100) / 100
	// success
	c.JSON(200, &response{
		Metal:    metal,
		Price:    price,
		Currency: "USD",
		Unit:     "oz",
	})

}
