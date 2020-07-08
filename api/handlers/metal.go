package handlers

import (
	"fmt"
	"strings"

	"github.com/chutified/metal-price/api/services"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPrice(c *gin.Context) {

	// get param
	m := c.Param("metal")
	m = strings.ToLower(m)

	// call the service
	price, err := services.GetPrice(h.mc, m)
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
