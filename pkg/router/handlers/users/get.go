package users_handlers

import (
	"app/pkg/logger"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetUsersHandler]")
	log.Debug("started")
	w.Write([]byte("GET ALL USERS"))

	log.Debug("finished")
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetUserByIDHandler]")
	log.Debug("started")
	w.Write([]byte("GET USER BY ID USERS"))

	log.Debug("finished")
}
