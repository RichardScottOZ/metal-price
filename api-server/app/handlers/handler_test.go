package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	services "github.com/chutified/metal-price/api-server/app/services"
	currency "github.com/chutified/metal-price/currency/service/protos/currency"
	metal "github.com/chutified/metal-price/metal/service/protos/metal"
	"google.golang.org/grpc"
)

func TestHandler(t *testing.T) {

	log := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	var h *Handler

	t.Run("NewHandler", func(t1 *testing.T) {

		currencyConn, err := grpc.Dial("localhsot:10551", grpc.WithInsecure())
		if err != nil {
			t.Fatalf("unable to dial localhost:10551: %v", err)
		}
		currencyClient := currency.NewCurrencyClient(currencyConn)
		cs := services.NewCurrency(currencyClient)

		metalConn, err := grpc.Dial("localhsot:10552", grpc.WithInsecure())
		if err != nil {
			t.Fatalf("unable to dial localhost:10521: %v", err)
		}
		metalClient := metal.NewMetalClient(metalConn)
		ms := services.NewMetal(metalClient)

		h = NewHandler(log, cs, ms)
	})

	t.Run("Ping", func(t1 *testing.T) {

		wri := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/ping", nil)

		// TODO
		_ = wri
		_ = req
		_ = err
		_ = h

	})
}
