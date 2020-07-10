package handlers

import "github.com/gin-gonic/gin"

// Ping determines whether there is a connection or not.
// @summary Check the status.
// @description Check if the server is running.
// @produce json
// @success 200 {object} handlers.Response "ok"
// @router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, &HTTPError{
		Message: "pong",
	})
}
