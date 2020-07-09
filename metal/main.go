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

	// init service
	s := service.NewService(logger)
	err = s.Init(cfg)
	if err != nil {
		logger.Fatalf("initialize the service: %v", err)
	}

	// run
	logger.Fatalf("run the service: %v", s.Run(cfg))
}
