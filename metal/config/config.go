package config

// Config holds the configuration values for metal microservice.
type Config struct {
	Port   int
	Source string
}

// GetConfig reads from the config file and returns the Config.
func GetConfig() *Config {
	return &cfg
}

var cfg = Config{
	Port:   10502,
	Source: "https://www.moneymetals.com/api/spot-prices.json",
}
