package config

// Config defines the app settings.
type Config struct {
	Port            int
	ExposePort      int
	CurrencyService string
	MetalService    string
	Debug           bool
}

// GetConfig returns the configuration.
func GetConfig() *Config {
	return &cfg
}

var cfg = Config{
	Port:            80,
	CurrencyService: "currencysrv:10501",
	MetalService:    "metalsrv:10502",
	Debug:           false,
}
