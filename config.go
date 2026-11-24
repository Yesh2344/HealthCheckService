package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config holds all configurable parameters for the service.
type Config struct {
	Port          string        `env:"PORT" envDefault:"8080"`
	Services      []string      `env:"SERVICES" envDefault:"http://localhost:8081,http://localhost:8082"`
	CheckInterval time.Duration `env:"CHECK_INTERVAL" envDefault:"30s"`
}

// Load reads environment variables (including a .env file if present) and populates Config.
func Load() (*Config, error) {
	// Load .env file; ignore error if file does not exist.
	_ = godotenv.Load()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}