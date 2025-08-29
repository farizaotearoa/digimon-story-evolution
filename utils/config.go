package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// AppConfig holds all application configuration
type AppConfig struct {
	App      AppInfo
	Logger   LoggerConfig
	Images   ImagesConfig
	CORS     CORSConfig
	Services ServicesConfig
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

// ServicesConfig holds configuration for all microservices
type ServicesConfig struct {
	DigimonService DigimonServiceConfig
	UserService    UserServiceConfig
	AdminService   AdminServiceConfig
}

type DigimonServiceConfig struct {
	BaseURL string
	Timeout time.Duration
}

type UserServiceConfig struct {
	BaseURL string
	Timeout time.Duration
}

type AdminServiceConfig struct {
	BaseURL string
	Timeout time.Duration
}

// Global config instance
var GlobalConfig *AppConfig

// Helper functions to get environment variables with defaults
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

func getEnvAsDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	} else {
		log.Println(".env file loaded successfully")
	}

	config := &AppConfig{
		App: AppInfo{
			Name: getEnvOrDefault("APP_NAME", "Digimon API Gateway"),
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
		Services: ServicesConfig{
			DigimonService: DigimonServiceConfig{
				BaseURL: getEnvOrDefault("DIGIMON_SERVICE_URL", "http://localhost:8081"),
				Timeout: getEnvAsDurationOrDefault("DIGIMON_SERVICE_TIMEOUT", 30*time.Second),
			},
			UserService: UserServiceConfig{
				BaseURL: getEnvOrDefault("USER_SERVICE_URL", "http://localhost:8082"),
				Timeout: getEnvAsDurationOrDefault("USER_SERVICE_TIMEOUT", 30*time.Second),
			},
			AdminService: AdminServiceConfig{
				BaseURL: getEnvOrDefault("ADMIN_SERVICE_URL", "http://localhost:8083"),
				Timeout: getEnvAsDurationOrDefault("ADMIN_SERVICE_TIMEOUT", 30*time.Second),
			},
		},
	}

	GlobalConfig = config
	return config
}
