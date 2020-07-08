package handlers

import "github.com/gin-gonic/gin"

// Ping determines whether there is a connection or not.
func Ping(c *gin.Context) {
	c.String(200, "pong")
}
