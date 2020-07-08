package main

import (
	"log"
	"os"

	"github.com/chutified/metal-price/api/handlers"
	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/chutified/metal-price/metal/protos/metal"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// currency client
	currencyConn, err := grpc.Dial("localhost:10501", grpc.WithInsecure())
	if err != nil {
		logger.Panicf("unable to dial: %v", err)
	}
	defer currencyConn.Close()
	cc := currency.NewCurrencyClient(currencyConn)

	// metal client
	metalConn, err := grpc.Dial("localhost:10502", grpc.WithInsecure())
	if err != nil {
		logger.Panicf("unable to dial: %v", err)
	}
	defer metalConn.Close()
	mc := metal.NewMetalClient(metalConn)

	// construct a new Handler
	h := handlers.NewHandler(logger, cc, mc)

	// gin router
	r := gin.New()

	// middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// routes
	r.GET("/ping", handlers.Ping)

	api := r.Group("/api")
	api.GET("/rate/:base/:dest", h.GetRate)
	api.GET("/metal/:metal", h.GetPrice)

	r.Run(":8080")
}
