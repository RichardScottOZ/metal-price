package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/chutified/metal-value-api/currency/protos/currency"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// gin router
	r := gin.New()
	r.Use(gin.Recovery())

	// currency client
	conn, err := grpc.Dial("localhost:10501")
	if err != nil {
		logger.Panic("Unable to dial: %v", err)
	}
	defer conn.Close()
	cc := currency.NewCurrencyClient(conn)

	// routes
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/rate/:base/:dest", func(c *gin.Context) {
		base, err := strconv.ParseInt(c.Param("base"), 10, 32)
		if err != nil {
			c.String(200, "invalid params")
			return
		}

		dest, err := strconv.ParseInt(c.Param("dest"), 10, 32)
		if err != nil {
			c.String(200, "invalid params")
			return
		}

		cc.GetRate(context.Background(), &currency.RateRequest{
			Base:        currency.Currencies(base),
			Destination: currency.Currencies(dest),
		})

		c.String(200, "pong")
	})

	r.GET("/")

	r.Run(":8080")
}
