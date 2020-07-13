package config

// Config holds the configuration values for currency microservice.
type Config struct {
	Port   int
	Source string
}

// GetConfig reads from the config file and returns the Config.
func GetConfig() *Config {
	return &cfg
}

var cfg = Config{
	Port:   10501,
	Source: "https://api.exchangeratesapi.io/latest",
}
