package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// gin router
	r := gin.New()
	r.Use(gin.Recovery())

	// currency client
	//currency.NewCurrencyClient()

	// routes
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/")

	r.Run(":9000")
}
