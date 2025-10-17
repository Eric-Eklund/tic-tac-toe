package config

import (
	"os"
)

type Config struct {
	Port  string
	Host  string
	Debug bool
}

func Load() *Config {
	return &Config{
		Port:  getEnv("PORT", "8080"),
		Host:  getEnv("HOST", "localhost"),
		Debug: getEnv("DEBUG", "false") == "true",
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
