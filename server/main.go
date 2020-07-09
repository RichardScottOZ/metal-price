package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chutified/metal-price/server/app"
	"github.com/chutified/metal-price/server/config"
)

func main() {

	// logging
	logger := log.New(os.Stdout, "metal-service: ", log.LstdFlags)

	// config
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		logger.Fatalf("get config: %v", err)
	}

	// init app
	a := app.NewApp()
	err = a.Init(cfg)
	if err != nil {
		panic(fmt.Errorf("initialize app: %w", err))
	}
	defer func() {
		errs := a.Stop()
		for i, err := range errs {
			fmt.Printf("close error %d: %v\n", i, err)
		}
	}()

	// run
	panic(fmt.Errorf("run app: %w", a.Run()))
}
