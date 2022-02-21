package config

import (
	"os"
)

type ServerConfig struct {
	ServerAddress string
	BaseURL       string
}

type FileStorageConfig struct {
	FileStoragePath string
}

type Config struct {
	Server      ServerConfig
	FileStorage FileStorageConfig
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
			BaseURL:       getEnv("BASE_URL", "http://localhost:8080"),
		},
		FileStorage: FileStorageConfig{
			FileStoragePath: getEnv("FILE_STORAGE_PATH", "urls.txt"),
		},
	}
}

func (config Config) GetBaseURL() string {
	return config.Server.BaseURL
}

func (config Config) GetServerAddress() string {
	return config.Server.ServerAddress
}

func (config Config) GetFileStoragePath() string {
	return config.FileStorage.FileStoragePath
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
