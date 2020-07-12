package config

// Config defines the app settings.
type Config struct {
	Port            int    `yaml:"PORT"`
	CurrencyService string `yaml:"Currency_Service"`
	MetalService    string `yaml:"Metal_Service"`
	Debug           bool   `yaml:"Debug_Mode"`
}

// GetConfig returns the configuration.
func GetConfig() *Config {
	return &cfg
}

var cfg = Config{
	Port:            8080,
	CurrencyService: "localhost:10501",
	MetalService:    "localhost:10502",
	Debug:           false,
}
