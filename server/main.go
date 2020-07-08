package main

import (
	"log"
	"os"

	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/chutified/metal-price/metal/protos/metal"
	"github.com/chutified/metal-price/server/handlers"
	"github.com/chutified/metal-price/server/services"
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
	currencyClient := currency.NewCurrencyClient(currencyConn)
	cs := services.NewCurrency(currencyClient)

	// metal client
	metalConn, err := grpc.Dial("localhost:10502", grpc.WithInsecure())
	if err != nil {
		logger.Panicf("unable to dial: %v", err)
	}
	defer metalConn.Close()
	metalClient := metal.NewMetalClient(metalConn)
	ms := services.NewMetal(metalClient)

	// construct a new Handler
	h := handlers.NewHandler(logger, cs, ms)

	// gin router
	r := gin.New()

	// middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// routes
	r.GET("/ping", handlers.Ping)

	api := r.Group("/api")
	api.GET("/:metal/:currency/*unit", h.GetPrice)

	r.Run(":8080")
}
