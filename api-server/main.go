package main

import (
	"log"
	"os"

	"github.com/chutified/metal-price/api-server/app"
	"github.com/chutified/metal-price/api-server/config"
)

func main() {
	logger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags)

	// config
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		logger.Fatalf("get config: %v", err)
	}

	// init app
	a := app.NewApp(logger)
	err = a.Init(cfg)
	if err != nil {
		logger.Panicf("initialize app: %v", err)
	}
	defer func() {
		errs := a.Stop()
		for i, err := range errs {
			logger.Printf("close error %d: %v\n", i, err)
		}
	}()

	// run
	logger.Panicf("run app: %v", a.Run())
}
