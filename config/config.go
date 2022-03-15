package config

import (
	"flag"
	"fmt"
	"os"
)

type ServerConfig struct {
	ServerAddress string
	BaseURL       string
}

type FileStorageConfig struct {
	FileStoragePath string
}

type EncryptionConfig struct {
	key string
}

type Config struct {
	Server      ServerConfig
	FileStorage FileStorageConfig
	Encryption  EncryptionConfig
}

func New() Config {
	config := Config{}
	config = initConfigByEnv(config)
	config = initConfigByFlag(config)

	return config
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

func (config Config) GetEncryptionKey() string {
	return config.Encryption.key
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func initConfigByEnv(config Config) Config {
	config.Server.ServerAddress = getEnv("SERVER_ADDRESS", "localhost:8080")
	config.Server.BaseURL = getEnv("BASE_URL", "http://localhost:8080")
	config.FileStorage.FileStoragePath = getEnv("FILE_STORAGE_PATH", "urls.txt")
	config.Encryption.key = getEnv("ENCRYPTION_KEY", "1234567890")

	return config
}

func initConfigByFlag(config Config) Config {
	if flag.Parsed() {
		fmt.Println("Error occurred. Re-initializing the config")
		return config
	}

	serverAddress := flag.String("a", "", "SERVER_ADDRESS")
	baseURL := flag.String("b", "", "BASE_URL")
	fileStoragePath := flag.String("f", "", "FILE_STORAGE_PATH")
	encryptionKey := flag.String("e", "", "ENCRYPTION_KEY")

	flag.Parse()

	if len(*serverAddress) > 0 {
		config.Server.ServerAddress = *serverAddress
	}

	if len(*baseURL) > 0 {
		config.Server.BaseURL = *baseURL
	}

	if len(*fileStoragePath) > 0 {
		config.FileStorage.FileStoragePath = *fileStoragePath
	}

	if len(*encryptionKey) > 0 {
		config.Encryption.key = *encryptionKey
	}

	return config
}
