package db

import (
	"app/pkg/config"
	"app/pkg/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var DB *sql.DB

func Connect(conf config.AppConfig) {
	db, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		logger.Log.Fatal("failed to connect to db", zap.String("DSN", conf.DSN), zap.Error(err))
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Fatal("failed to ping db", zap.String("DSN", conf.DSN), zap.Error(err))
	}

	DB = db
}
