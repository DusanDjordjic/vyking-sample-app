package players_handlers

import (
	"app/pkg/logger"
	"net/http"
)

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[CreatePlayerHandler]")
	log.Debug("started")
	w.Write([]byte("CREATE USER"))

	log.Debug("finished")
}
