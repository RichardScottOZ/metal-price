package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v2"
)

// Config holds the configuration values for metal microservice.
type Config struct {
	Port   int    `yaml:"PORT"`
	Source string `yaml:"source"`
}

// GetConfig reads from the config file and returns the Config.
func GetConfig(yamlPath string) (*Config, error) {

	// read file
	configPath := path.Join(rootDir(), yamlPath)
	bs, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	// get cfg
	var cfg Config
	err = yaml.UnmarshalStrict(bs, &cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal yaml file content: %w", err)
	}

	return &cfg, nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
