package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// MetalPrice return the price of the ounce of the metal in USD.
func (h *Handler) MetalPrice(c *gin.Context) {

	// get param
	m := c.Param("metal")
	m = strings.ToLower(m)

	// call the service
	price, err := h.ms.GetPrice(m)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("call metal service: %v", err),
		})
		return
	}

	// success
	c.JSON(200, gin.H{
		"price": price,
	})
}
