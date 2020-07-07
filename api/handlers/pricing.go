package handlers

import (
	"github.com/chutified/metal-price/api/services"
	"github.com/gin-gonic/gin"
)

// GetPrice returns price of a metal.
// TODO temporary
func GetPrice(c *gin.Context) {

	// get paramas
	base, dest := c.Param("base"), c.Param("dest")

	// call the service
	rate, err := services.GetRate(c, base, dest)
	if err != nil {
		c.String(500, "unable to call currency service: %v", err)
		return
	}

	// success
	c.JSON(200, gin.H{
		"rate": rate,
	})
}
