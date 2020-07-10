package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	config "github.com/chutified/metal-price/currency/config"
	"gopkg.in/go-playground/assert.v1"
)

func TestService(t *testing.T) {

	l := log.New(bytes.NewBufferString(""), "", log.LstdFlags)
	cfg := &config.Config{
		Port:   10551,
		Source: "https://api.exchangeratesapi.io/latest",
	}

	s := NewService(l, cfg)
	s.Init()

	assert.NotEqual(t, s.logger, nil)
	assert.NotEqual(t, s.srv, nil)

	tests := []struct {
		name      string
		action    func()
		expErrMsg string
	}{
		{
			name:      "ok",
			action:    func() { go http.ListenAndServe(":10551", nil) },
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
