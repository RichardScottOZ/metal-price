package main

import (
	"log"
	"os"

	app "github.com/chutified/metal-price/api-server/app"
	config "github.com/chutified/metal-price/api-server/config"
	_ "github.com/chutified/metal-price/api-server/docs" // documentation
)

// @title Metal Price API
// @version 1.0
// @description This API returns the current price of precious metals in different currencies and weight units.

// @contact.name Tommy Chu
// @contact.email tommychu2256@gmail.com

// @schemes http
// @host localhost:8080
// @BasePath /
func main() {
	logger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags)

	// config
	cfg := config.GetConfig()

	// init app
	a := app.NewApp(logger)
	err := a.Init(cfg)
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
