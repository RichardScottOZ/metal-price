package main

import (
	"log"
	"os"

	"github.com/chutified/metal-price/api/handlers"
	"github.com/chutified/metal-price/api/middlewares"
	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// currency client
	conn, err := grpc.Dial("localhost:10501", grpc.WithInsecure())
	if err != nil {
		logger.Panicf("unable to dial: %v", err)
	}
	defer conn.Close()
	cc := currency.NewCurrencyClient(conn)

	// gin router
	r := gin.New()

	// middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// routes
	r.GET("/ping", handlers.Ping)

	// pricing
	api := r.Group("/api")
	api.Use(middlewares.CurrencyClientMiddleware(cc))
	api.GET("/rate/:base/:dest", handlers.GetPrice) // TODO temp, refactor

	r.Run(":8080")
}
