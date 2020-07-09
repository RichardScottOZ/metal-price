package service

import (
	"fmt"
	"log"
	"net"

	"github.com/chutified/metal-price/metal/config"
	"github.com/chutified/metal-price/metal/service/data"
	"github.com/chutified/metal-price/metal/service/protos/metal"
	"github.com/chutified/metal-price/metal/service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Service is the service controller.
type Service struct {
	logger *log.Logger
	srv    *grpc.Server
}

// NewService constructs a new service controller.
func NewService(l *log.Logger) *Service {
	return &Service{
		logger: l,
	}
}

// Init defines service attributes.
func (s *Service) Init(cfg *config.Config) error {

	// data service
	pricesService, err := data.NewPrices(s.logger, cfg.Source)
	if err != nil {
		return fmt.Errorf("could not construct metal price data service: %w", err)
	}

	// servers
	metalServer := server.NewMetal(s.logger, pricesService)
	grpcSrv := grpc.NewServer()

	// register server
	metal.RegisterMetalServer(grpcSrv, metalServer)
	// reflection
	reflection.Register(grpcSrv)

	// success
	s.srv = grpcSrv
	return nil
}

// Run starts the service.
func (s *Service) Run(cfg *config.Config) error {

	// define listen
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("unable to listen: %w", err)
	}

	// listen
	s.logger.Printf("Listening gRPC on port %d", cfg.Port)
	return s.srv.Serve(listen)
}
