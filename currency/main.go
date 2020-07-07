package main

import (
	"log"
	"net"
	"os"

	"github.com/chutified/metal-price/currency/data"
	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/chutified/metal-price/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// data service
	rateService, err := data.NewRates(logger)
	if err != nil {
		logger.Fatalf("unable to construct data service: %v", err)
	}

	// server
	currencySrv := server.NewCurrency(logger, rateService)
	grpcSrv := grpc.NewServer()

	// register the server
	currency.RegisterCurrencyServer(grpcSrv, currencySrv)
	// register the reflection
	reflection.Register(grpcSrv)

	lst, err := net.Listen("tcp", ":10501")
	if err != nil {
		logger.Fatalf("unable to listen: %v", err)
	}
	grpcSrv.Serve(lst)
}
