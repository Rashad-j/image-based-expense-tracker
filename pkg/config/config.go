// Package config is the logic for the project configuration.
// it Utilizes the builder pattern
package config

import "github.com/caarlos0/env"

type Config struct {
	// server
	Port string `env:"PORT" envDefault:"8083"`
	// google
	GoogleApplicationCredentials string `env:"GOOGLE_APPLICATION_CREDENTIALS" envDefault:"./google-application-credentials.json"`
	// chatgpt
	ChatGPTApiKey string `env:"CHAT_GPT_API_KEY" envDefault:""`
	ChatGPTModel  string `env:"CHAT_GPT_MODEL" envDefault:"gpt-4o-mini"`
	// other
	LoggerLevel string `env:"LOGGER_LEVEL" envDefault:"0"`
}

func ReadEnvConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) WithPort(port string) *Config {
	c.Port = port
	return c
}

func (c *Config) WithGoogleApplicationCredentials(credentials string) *Config {
	c.GoogleApplicationCredentials = credentials
	return c
}

func (c *Config) WithChatGPTApiKey(key string) *Config {
	c.ChatGPTApiKey = key
	return c
}

func (c *Config) WithChatGPTModel(model string) *Config {
	c.ChatGPTModel = model
	return c
}
