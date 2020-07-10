package service

import (
	"fmt"
	"log"
	"net"

	config "github.com/chutified/metal-price/metal/config"
	metal "github.com/chutified/metal-price/metal/service/protos/metal"
	server "github.com/chutified/metal-price/metal/service/server"
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

	// servers
	metalServer := server.NewMetal(s.logger, cfg)
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
