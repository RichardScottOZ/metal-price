package handlers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	services "github.com/chutified/metal-price/api-server/app/services"
	currency "github.com/chutified/metal-price/currency/service/protos/currency"
	metal "github.com/chutified/metal-price/metal/service/protos/metal"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/assert.v1"
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

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	e := h.SetRoutes(gin.New())

	t.Run("Ping", func(t1 *testing.T) {

		w := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t1.Fatalf("unable to create new: %s", err.Error())
		}

		e.ServeHTTP(w, r)
		body, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t1.Fatalf("unable to read response body: %s", err.Error())
		}

		assert.Equal(t1, w.Code, 200)
		assert.Equal(t1, string(body), `{"message":"pong"}`)
	})

}
