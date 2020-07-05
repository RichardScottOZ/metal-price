package main

import (
	"github.com/chutified/metal-value-api/currency/protos/currency"
	"github.com/gin-gonic/gin"
)

func main() {

	// gin router
	r := gin.New()
	r.Use(gin.Recovery())

	// currency client
	currency.NewCurrencyClient()

	// routes
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":9000")
}
