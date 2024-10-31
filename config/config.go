package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	AppURL      string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		log.Fatal("DB_URL must be set")
	}

	appURL := os.Getenv("APP_URL")
	if appURL == "" {
		log.Fatal("APP_URL must be set")
	}

	return Config{
		Port:        port,
		DatabaseURL: databaseURL,
		AppURL:      appURL,
	}
}
