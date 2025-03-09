package main

import (
	"app/pkg/config"
	"app/pkg/db"
	"app/pkg/logger"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load envs, %s", err)
		os.Exit(1)
	}

	var migrationsFolder string
	flag.StringVar(&migrationsFolder, "migrationsFolder", "migrations", "specify port to listen on")
	flag.Parse()

	err = logger.Setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger %s\n", err)
		os.Exit(1)
	}

	appConfig := config.Parse()
	db.Connect(appConfig)
	logger.Log.Info("connected to database")

	entries, err := os.ReadDir(migrationsFolder)

	for _, e := range entries {
		if strings.HasSuffix(e.Name(), "up.sql") {
			filename := filepath.Join(migrationsFolder, e.Name())
			content, err := os.ReadFile(filename)
			if err != nil {
				logger.Log.Error("failed to read file", zap.String("Filename", filename), zap.Error(err))
				continue
			}

			_, err = db.DB.Exec(string(content))
			if err != nil {
				logger.Log.Error("failed to execure migration", zap.String("Filename", filename), zap.Error(err))
			} else {
				logger.Log.Info("execited migration", zap.String("Filename", filename))
			}
		}
	}

}
