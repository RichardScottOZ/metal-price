package middlewares

import (
	"github.com/chutified/metal-price/currency/protos/currency"
	"github.com/gin-gonic/gin"
)

// CUrrencyClientMiddleware defines currency for gin context.
func CurrencyClientMiddleware(cc currency.CurrencyClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("currency_client", cc)
		c.Next()
	}
}
