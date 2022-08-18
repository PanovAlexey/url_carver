// Package config - a package that stores the global application configuration with all settings.
// It is filled at the start of the application with data from the ENV file and the console.
package config

import (
	"flag"
	"fmt"
	"os"
)

type ServerConfig struct {
	ServerAddress string
	DebugAddress  string
	BaseURL       string
}

type FileStorageConfig struct {
	FileStoragePath string
}

type EncryptionConfig struct {
	key string
}

type DatabaseConfig struct {
	dsn string
}

type ApplicationConfig struct {
	IsDebug bool
}

type Config struct {
	Server      ServerConfig
	FileStorage FileStorageConfig
	Encryption  EncryptionConfig
	Database    DatabaseConfig
	Application ApplicationConfig
}

// New returns the initialized configuration structure
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

func (config Config) GetServerDebugAddress() string {
	return config.Server.DebugAddress
}

func (config Config) GetFileStoragePath() string {
	return config.FileStorage.FileStoragePath
}

func (config Config) GetEncryptionKey() string {
	return config.Encryption.key
}

func (config Config) GetDatabaseDsn() string {
	return config.Database.dsn
}

func (config Config) IsDebug() bool {
	return config.Application.IsDebug
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func initConfigByEnv(config Config) Config {
	config.Server.ServerAddress = getEnv("SERVER_ADDRESS", "localhost:8080")
	config.Server.DebugAddress = getEnv("DEBUG_ADDRESS", "0.0.0.0:8081")
	config.Server.BaseURL = getEnv("BASE_URL", "http://localhost:8080")
	config.FileStorage.FileStoragePath = getEnv("FILE_STORAGE_PATH", "urls.txt")
	config.Encryption.key = getEnv("ENCRYPTION_KEY", "1234567890")
	config.Database.dsn = getEnv("DATABASE_DSN", "postgresql://user_name:user_password@database_host:5432/database_name")
	config.Application.IsDebug = getEnv("IS_DEBUG", "false") == "true"

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
	databaseDsn := flag.String("d", "", "DATABASE_DSN")

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

	if len(*databaseDsn) > 0 {
		config.Database.dsn = *databaseDsn
	}

	return config
}
