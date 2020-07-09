package handlers

import (
	"log"

	services "github.com/chutified/metal-price/api-server/app/services"
)

// Handler defines the api handler.
type Handler struct {
	log *log.Logger
	cs  *services.Currency
	ms  *services.Metal
}

// NewHandler is the constructor of the Handler.
func NewHandler(l *log.Logger, cs *services.Currency, ms *services.Metal) *Handler {
	return &Handler{
		log: l,
		cs:  cs,
		ms:  ms,
	}
}
