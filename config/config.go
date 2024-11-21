package config

import "github.com/caarlos0/env"

type Config struct {
	Port string `env:"PORT" envDefault:":8083"`
}

func ReadEnvConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
