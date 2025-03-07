package config

import (
	"app/pkg/logger"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type AppConfig struct {
	DSN string
}

func Parse() AppConfig {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("failed to load env variables", zap.Error(err))
	}

	dsn := os.Getenv("DB_DSN")
	return AppConfig{
		DSN: dsn,
	}
}
