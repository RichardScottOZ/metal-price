package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"

	config "github.com/chutommy/metal-price/currency/config"
	currency "github.com/chutommy/metal-price/currency/service/protos/currency"
	"gopkg.in/go-playground/assert.v1"
)

func TestNewCurrency(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	cfg := &config.Config{
		Port:   10551,
		Source: "https://api.exchangeratesapi.io/latest",
	}

	m := NewCurrency(l, cfg)

	assert.NotEqual(t, m.log, nil)
	assert.NotEqual(t, m.cfg, nil)
}

func TestGetRate(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	cfg := &config.Config{
		Port: 10551,
	}

	tests := []struct {
		name      string
		source    string
		baseID    int32
		destID    int32
		expErrMsg string
	}{
		{
			name:      "ok",
			source:    "https://api.exchangeratesapi.io/latest",
			baseID:    0,
			destID:    0,
			expErrMsg: "",
		},
		{
			name:      "invalid source",
			source:    "https://www.test.com",
			baseID:    0,
			destID:    0,
			expErrMsg: "could not construct currency price data service",
		},
		{
			name:      "invalid metal",
			source:    "https://api.exchangeratesapi.io/latest",
			baseID:    1052019,
			destID:    1052019,
			expErrMsg: "could not call GetRate currency service",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			cfg.Source = test.source
			c := NewCurrency(l, cfg)

			req := &currency.RateRequest{
				Base:        currency.Currencies(test.baseID),
				Destination: currency.Currencies(test.destID),
			}
			resp, err := c.GetRate(context.Background(), req)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.NotEqual(t1, resp.GetRate(), 1)
				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
