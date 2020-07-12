package main

import (
	"log"
	"os"

	config "github.com/chutified/metal-price/metal/config"
	service "github.com/chutified/metal-price/metal/service"
)

func main() {
	logger := log.New(os.Stdout, "[METAL SERVICE] ", log.LstdFlags)

	// config
	cfg := config.GetConfig()

	// service
	s := service.NewService(logger, cfg)
	s.Init()
	logger.Fatalf("run the service: %v", s.Run())
}
