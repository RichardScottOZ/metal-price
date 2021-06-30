package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"testing"

	config "github.com/chutommy/metal-price/metal/config"
	metal "github.com/chutommy/metal-price/metal/service/protos/metal"
	"gopkg.in/go-playground/assert.v1"
)

func TestNewMetal(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	cfg := &config.Config{
		Port:   10552,
		Source: "https://www.moneymetals.com/api/spot-prices.json",
	}

	m := NewMetal(l, cfg)

	assert.NotEqual(t, m.log, nil)
	assert.NotEqual(t, m.cfg, nil)
}

func TestGetPrice(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	cfg := &config.Config{
		Port: 10552,
	}

	tests := []struct {
		name      string
		source    string
		metalID   int32
		expErrMsg string
	}{
		{
			name:      "ok",
			source:    "https://www.moneymetals.com/api/spot-prices.json",
			metalID:   0,
			expErrMsg: "",
		},
		{
			name:      "invalid source",
			source:    "https://www.test.com",
			metalID:   0,
			expErrMsg: "could not construct metal price data service",
		},
		{
			name:      "invalid metal",
			source:    "https://www.moneymetals.com/api/spot-prices.json",
			metalID:   1052019,
			expErrMsg: "unable to get the price of the material",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			cfg.Source = test.source
			m := NewMetal(l, cfg)

			req := &metal.MetalRequest{
				Metal: metal.Materials(test.metalID),
			}
			resp, err := m.GetPrice(context.Background(), req)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.NotEqual(t1, resp.GetPrice(), 0)
				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
