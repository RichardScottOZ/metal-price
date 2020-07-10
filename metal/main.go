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
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		logger.Fatalf("get config: %v", err)
	}

	// service
	s := service.NewService(logger, cfg)
	s.Init()
	logger.Fatalf("run the service: %v", s.Run())
}
