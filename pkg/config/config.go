package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DSN string
}

func Parse() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env variables, error: %s", err)
	}

	dsn := os.Getenv("DB_DSN")
	return AppConfig{
		DSN: dsn,
	}
}
