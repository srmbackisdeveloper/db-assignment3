package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	LogLevel      string
	Port  string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://postgres:XoygASBVDxiPvbgYSfVyKCSfJejIWCJn@autorack.proxy.rlwy.net:11058/railway"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
		Port: getEnv("PORT", "8080"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}