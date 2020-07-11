package handlers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetRoutes set the app engine and its routing.
func (h *Handler) SetRoutes(e *gin.Engine) *gin.Engine {
	e = gin.New()

	// middlewares
	e.Use(gin.Recovery())
	e.Use(gin.Logger())

	// routes
	e.GET("/ping", Ping)

	api := e.Group("/i")
	api.GET("/:metal/:currency/:unit", h.GetMetalMCU)
	api.GET("/:metal/:currency", h.GetMetalMC)
	api.GET("/:metal", h.GetMetalM)

	// documentation
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// done
	return e
}
