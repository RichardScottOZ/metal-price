package config

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {

	tests := []struct {
		name      string
		yaml      string
		expPort   int
		expSource string
		expErrMsg string
	}{
		{
			name:      "ok",
			yaml:      "config/test/test_0.yaml",
			expPort:   10552,
			expSource: "https://test.com/metal.json",
			expErrMsg: "",
		},
		{
			name:      "invalid file",
			yaml:      "config/test/test_1.yaml",
			expPort:   0,
			expSource: "",
			expErrMsg: "unmarshal yaml file content",
		},
		{
			name:      "non-existing file",
			yaml:      "config/test/test_2.yaml",
			expPort:   0,
			expSource: "",
			expErrMsg: "open file",
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
				assert.Equal(t1, cfg.Source, test.expSource)
				assert.Equal(t1, "", test.expErrMsg)
			}
		})
	}
}
