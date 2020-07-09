package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/chutified/metal-price/metal/protos/metal"
	"github.com/chutified/metal-price/server/app/handlers"
	"github.com/chutified/metal-price/server/app/services"
	"github.com/chutified/metal-price/server/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// App is a server controller.
type App struct {
	logger      *log.Logger
	server      *http.Server
	engine      *gin.Engine
	connections []*grpc.ClientConn
}

// NewApp returns new App controller.
func NewApp(l *log.Logger) *App {
	return &App{
		logger: l,
	}
}

// Init sets everything for the App controller.
func (a *App) Init(cfg *config.Config) error {

	// currency client
	currencyConn, err := grpc.Dial(cfg.CurrencyService, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial: %w", err)
	}
	a.connections = append(a.connections, currencyConn) // CONN
	cs := services.NewCurrency(currency.NewCurrencyClient(currencyConn))

	// metal client
	metalConn, err := grpc.Dial(cfg.MetalService, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to dial: %w", err)
	}
	a.connections = append(a.connections, metalConn) // CONN
	ms := services.NewMetal(metal.NewMetalClient(metalConn))

	// construct an engine
	handler := handlers.NewHandler(a.logger, cs, ms)
	a.SetRoutes(handler) // ENGINE

	// server
	a.server = &http.Server{ // SERVER
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           a.engine,
		ReadTimeout:       4 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      4 * time.Second,
		IdleTimeout:       6 * time.Second,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	// success
	return nil
}

// Stop gracefully close all service's connections.
func (a *App) Stop() []error {
	var errs []error
	for _, conn := range a.connections {
		errs = append(errs, conn.Close())
	}
	return errs
}

// Run starts the server.
func (a *App) Run() error {
	a.logger.Printf("Listening and serving HTTP on port %s.\n", a.server.Addr)
	return a.server.ListenAndServe()
}
