package handlers

import (
	"fmt"
	"strings"

	"github.com/chutified/metal-price/api/services"
	"github.com/gin-gonic/gin"
)

// GetPrice returns price of a metal.
func (h *Handler) GetRate(c *gin.Context) {

	// get paramas
	base, dest := c.Param("base"), c.Param("dest")
	base, dest = strings.ToUpper(base), strings.ToUpper(dest)

	// call the service
	rate, err := services.GetRate(h.cc, base, dest)
	if err != nil {
		c.JSON(400, gin.H{
			"error": fmt.Sprintf("unable to call currency service: %v", err),
		})
		return
	}

	// success
	c.JSON(200, gin.H{
		"rate": rate,
	})
}
