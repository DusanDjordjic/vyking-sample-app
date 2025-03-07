package users_handlers

import (
	"app/pkg/logger"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[CreateUserHandler]")
	log.Debug("started")
	w.Write([]byte("CREATE USER"))

	log.Debug("finished")
}
