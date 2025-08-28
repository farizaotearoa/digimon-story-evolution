package utils

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Database DatabaseConfig
	App      AppInfo
	Logger   LoggerConfig
	Images   ImagesConfig
	CORS     CORSConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	SSLMode  string
}

type AppInfo struct {
	Name string
	Port string
}

type LoggerConfig struct {
	Path string
}

type ImagesConfig struct {
	Path string
}

type CORSConfig struct {
	AllowOrigins []string
}

var GlobalConfig *AppConfig

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsStringSliceOrDefault(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}

func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	} else {
		log.Println(".env file loaded successfully")
	}

	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("APP_NAME: %s", os.Getenv("APP_NAME"))

	config := &AppConfig{
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvAsIntOrDefault("DB_PORT", 5432),
			Name:     getEnvOrDefault("DB_NAME", "database"),
			Username: getEnvOrDefault("DB_USERNAME", ""),
			Password: getEnvOrDefault("DB_PASSWORD", ""),
			SSLMode:  getEnvOrDefault("DB_SSL_MODE", "disable"),
		},
		App: AppInfo{
			Name: getEnvOrDefault("APP_NAME", "Digimon Story Wiki"),
			Port: getEnvOrDefault("APP_PORT", "8080"),
		},
		Logger: LoggerConfig{
			Path: getEnvOrDefault("LOGGER_PATH", "./logs"),
		},
		Images: ImagesConfig{
			Path: getEnvOrDefault("IMAGES_PATH", "./images"),
		},
		CORS: CORSConfig{
			AllowOrigins: getEnvAsStringSliceOrDefault("CORS_ALLOW_ORIGINS", []string{"http://localhost:3000"}),
		},
	}

	GlobalConfig = config
	return config
}
