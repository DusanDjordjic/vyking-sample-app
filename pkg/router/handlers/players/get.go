package players_handlers

import (
	"app/pkg/config"
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/models"
	"app/pkg/router/response"
	"app/pkg/services"
	"app/pkg/utils"
	"net/http"

	"go.uber.org/zap"
)

func GetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetPlayersHandler]")
	log.Debug("started")

	limit := utils.GetInt64QueryParamWithDefault(r, "limit", config.DEFAULT_LIMIT)
	limit = utils.ValidateLimit(limit)
	log = log.With(zap.Int64("Limit", limit))

	offset := utils.GetInt64QueryParamWithDefault(r, "offset", 0)
	offset = utils.ValidateOffset(offset)
	log = log.With(zap.Int64("Offset", offset))

	players, err := services.PlayersGet(db.DB, limit, offset)
	if err != nil {
		response.NewInternalError(w)
		return
	}

	log = log.With(zap.Int("Count", len(players)))

	dtos := models.PlayerDTOs(players)
	err = response.JSONResponse(w, http.StatusOK, dtos)
	if err != nil {
		log.Error("failed to send players response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}

func GetPlayerByIDHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetPlayerByIDHandler]")
	log.Debug("started")

	// /api/users/{id}
	playerID, err := utils.GetInt64PathParameter(r.URL, 3)
	if err != nil {
		log.Warn("failed to get user id", zap.Error(err))
		response.NewBadRequest(w)
		return
	}
	log = log.With(zap.Int64("PlayerID", playerID))

	player, err := services.PlayersGetByID(db.DB, playerID)
	if err != nil {
		log.Warn("failed to get player by id", zap.Error(err))
		// FIXME: parse the error and return 500 if something is wrong with database
		response.NewNotFound(w)
		return
	}

	err = response.JSONResponse(w, http.StatusOK, player.ToDTO())
	if err != nil {
		log.Error("failed to send players response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
