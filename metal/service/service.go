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
	log *log.Logger
	srv *grpc.Server
	cfg *config.Config
}

// NewService constructs a new service controller.
func NewService(l *log.Logger, cfg *config.Config) *Service {
	return &Service{
		log: l,
		cfg: cfg,
	}
}

// Init defines service attributes.
func (s *Service) Init() {

	// servers
	metalServer := server.NewMetal(s.log, s.cfg)
	grpcSrv := grpc.NewServer()

	// register server
	metal.RegisterMetalServer(grpcSrv, metalServer)
	// reflection
	reflection.Register(grpcSrv)

	// success
	s.srv = grpcSrv
}

// Run starts the service.
func (s *Service) Run() error {

	// define listen
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return fmt.Errorf("unable to listen: %w", err)
	}

	// listen
	s.log.Printf("Metal service is running (active)")
	return s.srv.Serve(listen)
}
