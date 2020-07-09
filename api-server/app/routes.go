package app

import (
	"github.com/chutified/metal-price/api-server/app/handlers"
	"github.com/gin-gonic/gin"
)

// SetRoutes set the app engine and its routing.
func (a *App) SetRoutes(h *handlers.Handler) {
	a.engine = gin.New()

	// middlewares
	a.engine.Use(gin.Recovery())
	a.engine.Use(gin.Logger())

	// routes
	a.engine.GET("/ping", handlers.Ping)

	api := a.engine.Group("/i")
	api.GET("/:metal/:currency/:unit", h.GetPrice)
}
