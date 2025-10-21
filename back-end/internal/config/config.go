package config

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

func GetCors() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

}
