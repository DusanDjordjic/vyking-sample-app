package main

import (
	"app/pkg/config"
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/router"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load envs, %s", err)
		os.Exit(1)
	}

	var port uint
	flag.UintVar(&port, "port", 8080, "specify port to listen on")
	flag.Parse()

	err = logger.Setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger %s\n", err)
		os.Exit(1)
	}

	appConfig := config.Parse()
	db.Connect(appConfig)
	logger.Log.Info("connected to database")

	mux := http.NewServeMux()
	router.SetupRouter(mux)

	logger.Log.Info("running server", zap.Uint("Port", port))
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		logger.Log.Error("failed to start server", zap.Uint("Port", port), zap.Error(err))
		os.Exit(1)
	}
}
