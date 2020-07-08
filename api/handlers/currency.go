package handlers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Exchange returns the exchage rate between the base and destination currency.
func (h *Handler) Exchange(c *gin.Context) {

	// get paramas
	base, dest := c.Param("base"), c.Param("dest")
	base, dest = strings.ToUpper(base), strings.ToUpper(dest)

	// call the service
	rate, err := h.cs.GetRate(base, dest)
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
