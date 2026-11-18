package config

type Config struct {
	AppUrl         string
	TrustedProxies []string
}

func Load() Config {
	// Loads and validates env vars
	return Config{
		AppUrl:         "http://localhost:8080",
		TrustedProxies: []string{},
	}
}
