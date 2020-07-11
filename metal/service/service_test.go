package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	config "github.com/chutified/metal-price/metal/config"
	"gopkg.in/go-playground/assert.v1"
)

func TestService(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	cfg := &config.Config{
		Port:   10552,
		Source: "https://www.moneymetals.com/api/spot-prices.json",
	}

	// NewService
	s := NewService(l, cfg)
	// Init
	s.Init()

	assert.NotEqual(t, s.log, nil)
	assert.NotEqual(t, s.srv, nil)

	tests := []struct {
		name      string
		action    func()
		expErrMsg string
	}{
		{
			name:      "ok",
			action:    func() { go http.ListenAndServe(":10552", nil) },
			expErrMsg: "",
		},
		{
			name:      "address already in use",
			action:    func() {},
			expErrMsg: "unable to listen",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			// Run
			test.action()
			var err error
			go func() {
				err = s.Run()
			}()
			time.Sleep(1200 * time.Millisecond)

			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
