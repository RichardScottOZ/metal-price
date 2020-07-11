package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	// >>>>>>>>>>>>>>> NewHandler
	var h *Handler
	var currencyConn, metalConn *grpc.ClientConn
	t.Run("NewHandler", func(t1 *testing.T) {
		var err error

		currencyConn, err = grpc.Dial("localhost:10501", grpc.WithInsecure())
		if err != nil {
			t.Fatalf("unable to dial localhost:10501: %v", err)
		}
		currencyClient := currency.NewCurrencyClient(currencyConn)
		cs := services.NewCurrency(currencyClient)

		metalConn, err = grpc.Dial("localhost:10502", grpc.WithInsecure())
		if err != nil {
			t.Fatalf("unable to dial localhost:10501: %v", err)
		}
		metalClient := metal.NewMetalClient(metalConn)
		ms := services.NewMetal(metalClient)

		h = NewHandler(log, cs, ms)
	})
	defer currencyConn.Close()
	defer metalConn.Close()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// >>>>>>>>>>>>>>> SetRoutes
	e := h.SetRoutes(gin.New())

	// >>>>>>>>>>>>>>> Ping
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

	tests := []struct {
		name      string
		metal     string
		currency  string
		unit      string
		expCode   int
		expErrMsg string
	}{
		// >>>>>>>>>>>>>>> GetMetalM
		{
			name:      "M ok",
			metal:     "silver",
			expCode:   200,
			expErrMsg: "",
		},
		{
			name:      "M invalid metal",
			metal:     "invalid",
			expCode:   400,
			expErrMsg: "call metal service",
		},
		// >>>>>>>>>>>>>>> GetMetalMC
		{
			name:      "MC ok",
			metal:     "silver",
			currency:  "EUR",
			expCode:   200,
			expErrMsg: "",
		},
		{
			name:      "MC invalid metal",
			metal:     "silver",
			currency:  "EUR",
			expCode:   400,
			expErrMsg: "call metal service",
		},
		{
			name:      "MC invalid currency",
			metal:     "silver",
			currency:  "invalid",
			expCode:   400,
			expErrMsg: "call currency service",
		},
		// >>>>>>>>>>>>>>> GetMetalMCU
		{
			name:      "MCU ok",
			metal:     "silver",
			currency:  "EUR",
			unit:      "lb",
			expCode:   200,
			expErrMsg: "",
		},
		{
			name:      "MCU invalid metal",
			metal:     "invalid",
			currency:  "EUR",
			unit:      "lb",
			expCode:   400,
			expErrMsg: "call metal service",
		},
		{
			name:      "MCU invalid currency",
			metal:     "silver",
			currency:  "invalid",
			unit:      "lb",
			expCode:   400,
			expErrMsg: "call currency service",
		},
		{
			name:      "MCU invalid unit",
			metal:     "silver",
			currency:  "EUR",
			unit:      "invalid",
			expCode:   400,
			expErrMsg: "call weight unit converter",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			path := fmt.Sprintf("/i/%s", test.metal)
			if test.currency != "" {
				path = fmt.Sprintf("%s/%s", path, test.currency)

				if test.unit != "" {
					path = fmt.Sprintf("%s/%s", path, test.unit)
				}
			}

			w := httptest.NewRecorder()
			r, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t1.Fatalf("unable to create new: %s", err.Error())
			}

			e.ServeHTTP(w, r)
			if err != nil {
				t1.Fatalf("unable to read response body: %s", err.Error())
			}

			fmt.Println(test.name, string(w.Body.Bytes()))

			if w.Code == 400 {

				var httpErr HTTPError
				err := json.NewDecoder(r.Body).Decode(&httpErr)
				if err != nil {
					t1.Fatalf("invalid error response: %s", err.Error())
				}

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, httpErr.Message, exp)
				assert.Equal(t1, 400, test.expCode)

			} else if w.Code == 200 {

				var resp Response
				err := json.NewDecoder(r.Body).Decode(&resp)
				if err != nil {
					t1.Fatalf("invalid response: %s", err.Error())
				}

				assert.Equal(t1, resp.Metal, test.metal)
				if test.currency != "" {
					assert.Equal(t1, resp.Currency, test.currency)

					if test.unit != "" {
						assert.Equal(t1, resp.Unit, test.unit)
					}
				}

				assert.Equal(t1, 200, test.expCode)
				assert.Equal(t1, "", test.expErrMsg)

			} else {

				t.Fatalf("unexpected error message: %d", w.Code)

			}
		})
	}
}
