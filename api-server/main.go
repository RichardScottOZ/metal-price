package main

import (
	"log"
	"os"

	app "github.com/chutified/metal-price/api-server/app"
	config "github.com/chutified/metal-price/api-server/config"
)

// @title Bookstore API example with Gin
// @version 1.0
// @description This is a sample of a Gin API framework.

// @contact.name Tommy Chu
// @contact.email tommychu2256@gmail.com

// @schemes http
// @host localhost:8081
// @BasePath /api/v1
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
