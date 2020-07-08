package main

import (
	"log"
	"net"
	"os"

	"github.com/chutified/metal-price/metal/data"
	"github.com/chutified/metal-price/metal/protos/metal"
	"github.com/chutified/metal-price/metal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.New(os.Stdout, "metal-service: ", log.LstdFlags)

	// data service
	pricesService, err := data.NewPrices(logger)
	if err != nil {
		logger.Fatalf("could not construct metal price data service: %v", err)
	}

	// server
	metalServer := server.NewMetal(logger, pricesService)
	grpcSrv := grpc.NewServer()

	// register server
	metal.RegisterMetalServer(grpcSrv, metalServer)
	// reflection
	reflection.Register(grpcSrv)

	// run server
	listen, err := net.Listen("tcp", ":10502")
	if err != nil {
		logger.Fatalf("unable to listen: %v", err)
	}
	grpcSrv.Serve(listen)
}
