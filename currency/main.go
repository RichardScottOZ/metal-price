package main

import (
	"log"
	"net"
	"os"

	"github.com/chutified/metal-value-api/currency/protos/currency"
	"github.com/chutified/metal-value-api/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	currencySrv := server.New(logger)
	grpcSrv := grpc.NewServer()

	currency.RegisterCurrencyServer(grpcSrv, currencySrv)
	reflection.Register(grpcSrv)

	lst, err := net.Listen("tcp", ":9091")
	if err != nil {
		logger.Fatalf("Unable to listen: %v", err)
	}
	grpcSrv.Serve(lst)
}
