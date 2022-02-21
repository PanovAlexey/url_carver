package config

import (
	"os"
)

type ServerConfig struct {
	ServerAddress string
	ServerPort    string
	BaseURL       string
}

type Config struct {
	Server ServerConfig
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			ServerAddress: getEnv("SERVER_ADDRESS", "localhost"),
			ServerPort:    getEnv("SERVER_PORT", "8080"),
			BaseURL:       getEnv("BASE_URL", "http://localhost"),
		},
	}
}

func (config Config) GetBaseURL() string {
	return config.Server.BaseURL
}

func (config Config) GetPort() string {
	return config.Server.ServerPort
}

func (config Config) GetServerAddress() string {
	return config.Server.ServerAddress
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
