package app

import (
	"bytes"
	"log"
	"testing"
	"time"

	config "github.com/chutommy/metal-price/api-server/config"
	"gopkg.in/go-playground/assert.v1"
)

func TestApp(t *testing.T) {

	logger := log.New(bytes.NewBufferString(""), "", log.LstdFlags)

	cfg := &config.Config{
		Port:            8080,
		CurrencyService: "localhost:10501",
		MetalService:    "localhost:10502",
		Debug:           false,
	}

	// >>>>>>>>>>>>>>> NewApp
	a := NewApp(logger)
	// >>>>>>>>>>>>>>> Init
	err := a.Init(cfg)
	if err != nil {

		t.Fatalf("unexpected error: %s", err.Error())

	}

	errs := a.Stop()
	assert.Equal(t, errs[0], nil)
	assert.Equal(t, errs[1], nil)

	assert.NotEqual(t, a.log, nil)
	assert.NotEqual(t, a.server, nil)
	assert.NotEqual(t, a.engine, nil)
	assert.Equal(t, len(a.connections), 2)

	err = nil
	// >>>>>>>>>>>>>>> Run
	go func() {
		err = a.Run()
		if err != nil {
			t.Fatalf("unexpected error: %s", err.Error())
		}
	}()
	time.Sleep(1500 * time.Millisecond)

	assert.Equal(t, err, nil)
}
