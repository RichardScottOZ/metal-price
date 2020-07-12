package main

import (
	"log"
	"os"

	config "github.com/chutified/metal-price/currency/config"
	service "github.com/chutified/metal-price/currency/service"
)

func main() {
	logger := log.New(os.Stdout, "[CURRENCY SERVICE] ", log.LstdFlags)

	// config
	cfg := config.GetConfig()

	// init service
	s := service.NewService(logger, cfg)
	s.Init()

	// run
	logger.Fatalf("run the service: %v", s.Run())
}
