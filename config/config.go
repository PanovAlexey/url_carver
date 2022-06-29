// Package config - a package that stores the global application configuration with all settings.
// It is filled at the start of the application with data from the ENV file and the console.
package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type ServerConfig struct {
	ServerAddress            string
	DebugAddress             string
	BaseURL                  string
	TimeoutShutdownInSeconds int
	EnableHTTPS              *bool
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
	IsDebug        *bool
	ConfigJSONPath string
}

type Config struct {
	Server      ServerConfig
	FileStorage FileStorageConfig
	Encryption  EncryptionConfig
	Database    DatabaseConfig
	Application ApplicationConfig
}

type ConfigJSON struct {
	ServerAddress   string `json:"server_address"`
	BaseURL         string `json:"base_url"`
	FileStoragePath string `json:"file_storage_path"`
	DatabaseDSN     string `json:"database_dsn"`
	EnableHTTPS     *bool  `json:"enable_https"`
}

// New returns the initialized configuration structure
func New() Config {
	config := Config{}
	config = initConfigByEnv(config)
	config = initConfigByFlag(config)
	config = initConfigByJSONConfig(config)
	config = initConfigByDefault(config)

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

func (config Config) GetConfigJSONPath() string {
	return config.Application.ConfigJSONPath
}

func (config Config) IsDebug() bool {
	return *config.Application.IsDebug
}

func (config Config) IsEnableHTTPS() bool {
	return *config.Server.EnableHTTPS
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}

func initConfigByEnv(config Config) Config {
	config.Server.ServerAddress = getEnv("SERVER_ADDRESS")
	config.Server.DebugAddress = getEnv("DEBUG_ADDRESS")
	config.Server.BaseURL = getEnv("BASE_URL")

	if len(getEnv("ENABLE_HTTPS")) > 0 {
		isEnableHTTPS := getEnv("ENABLE_HTTPS") == "true"
		config.Server.EnableHTTPS = &isEnableHTTPS
	}

	config.FileStorage.FileStoragePath = getEnv("FILE_STORAGE_PATH")
	config.Encryption.key = getEnv("ENCRYPTION_KEY")
	config.Database.dsn = getEnv("DATABASE_DSN")

	config.Application.ConfigJSONPath = getEnv("CONFIG")

	if len(getEnv("IS_DEBUG")) > 0 {
		isDebug := getEnv("IS_DEBUG") == "true"
		config.Application.IsDebug = &isDebug
	}

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
	shortConfigJSONPath := flag.String("c", "", "CONFIG")
	longConfigJSONPath := flag.String("config", "", "CONFIG")

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

	if len(*shortConfigJSONPath) > 0 {
		config.Application.ConfigJSONPath = *shortConfigJSONPath
	}

	if len(*longConfigJSONPath) > 0 {
		config.Application.ConfigJSONPath = *longConfigJSONPath
	}

	return config
}

func initConfigByJSONConfig(config Config) Config {
	if len(config.Application.ConfigJSONPath) < 1 {
		return config
	}

	jsonConfigFile, err := os.OpenFile(config.Application.ConfigJSONPath, os.O_RDONLY, 0644)
	configJSON := ConfigJSON{}

	if err != nil {
		log.Println("error reading json config file " + err.Error())
	} else {
		jsonBody, err := io.ReadAll(jsonConfigFile)

		if err != nil {
			log.Println("error reading json config file " + err.Error())
		} else if err = json.Unmarshal(jsonBody, &configJSON); err != nil {
			log.Println("error reading json config file " + err.Error())
		}
	}

	if config.Server.EnableHTTPS == nil {
		config.Server.EnableHTTPS = configJSON.EnableHTTPS
	}

	if len(config.FileStorage.FileStoragePath) < 1 && len(configJSON.FileStoragePath) > 0 {
		config.FileStorage.FileStoragePath = configJSON.FileStoragePath
	}

	if len(config.Server.ServerAddress) < 1 && len(configJSON.ServerAddress) > 0 {
		config.Server.ServerAddress = configJSON.ServerAddress
	}

	if len(config.Server.BaseURL) < 1 && len(configJSON.BaseURL) > 0 {
		config.Server.BaseURL = configJSON.BaseURL
	}

	if len(config.Database.dsn) < 1 && len(configJSON.DatabaseDSN) > 0 {
		config.Database.dsn = configJSON.DatabaseDSN
	}

	return config
}

func initConfigByDefault(config Config) Config {
	if len(config.Server.ServerAddress) < 1 {
		config.Server.ServerAddress = "localhost:8080"
	}

	if len(config.Server.DebugAddress) < 1 {
		config.Server.DebugAddress = "localhost:8081"
	}

	if len(config.Server.BaseURL) < 1 {
		config.Server.BaseURL = "http://0.0.0.0:8080"
	}

	if config.Server.TimeoutShutdownInSeconds < 1 {
		config.Server.TimeoutShutdownInSeconds = 15
	}

	if config.Server.EnableHTTPS == nil {
		defaultEnableHTTPS := false
		config.Server.EnableHTTPS = &defaultEnableHTTPS
	}

	if len(config.FileStorage.FileStoragePath) < 1 {
		config.FileStorage.FileStoragePath = "urls.txt"
	}

	if len(config.Encryption.key) < 1 {
		config.Encryption.key = "234324324324234324234"
	}

	if len(config.Database.dsn) < 1 {
		config.Database.dsn = "postgresql://postgresql_user:user_password@postgres_container:5432/postgresql"
	}

	if config.Application.IsDebug == nil {
		defaultIsDebug := false
		config.Application.IsDebug = &defaultIsDebug
	}

	return config
}
