package config

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {

	cfg := GetConfig()

	assert.NotEqual(t, cfg.Port, 0)
	assert.NotEqual(t, cfg.Source, "")
}
