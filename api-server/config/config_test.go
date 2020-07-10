package config

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {

	tests := []struct {
		name            string
		yaml            string
		expPort         int
		expCurrencyHost string
		expMetalHost    string
		expDebugMode    bool
		expErrMsg       string
	}{
		{
			name:            "ok",
			yaml:            "config/test/test_0.yaml",
			expPort:         8081,
			expCurrencyHost: "localhost:10551",
			expMetalHost:    "localhost:10552",
			expDebugMode:    false,
			expErrMsg:       "",
		},
		{
			name:            "invalid file",
			yaml:            "config/test/test_1.yaml",
			expPort:         0,
			expCurrencyHost: "",
			expMetalHost:    "",
			expDebugMode:    false,
			expErrMsg:       "unmarshal yaml file content",
		},
		{
			name:            "non-existing file",
			yaml:            "config/test/test_2.yaml",
			expPort:         0,
			expCurrencyHost: "",
			expMetalHost:    "",
			expDebugMode:    false,
			expErrMsg:       "open file",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			cfg, err := GetConfig(test.yaml)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, cfg.Port, test.expPort)
				assert.Equal(t1, cfg.CurrencyService, test.expCurrencyHost)
				assert.Equal(t1, cfg.MetalService, test.expMetalHost)
				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
