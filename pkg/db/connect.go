package db

import (
	"app/pkg/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(conf config.AppConfig) {
	db, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		log.Fatalf("failed to connect to db dsn: %s, error: %s\n", conf.DSN, err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping db, error: %s", err)
	}

	DB = db
}
