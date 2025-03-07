package tournaments_handlers

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

func GetTournamentsHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetTournamentsHandler]")
	log.Debug("started")

	limit := utils.GetInt64QueryParamWithDefault(r, "limit", config.DEFAULT_LIMIT)
	limit = utils.ValidateLimit(limit)
	log = log.With(zap.Int64("Limit", limit))

	offset := utils.GetInt64QueryParamWithDefault(r, "offset", 0)
	offset = utils.ValidateOffset(offset)
	log = log.With(zap.Int64("Offset", offset))

	tournaments, err := services.TournamentsGet(db.DB, limit, offset)
	if err != nil {
		log.Error("failed to get tournaments", zap.Error(err))
		response.NewInternalError(w)
		return
	}

	log = log.With(zap.Int("Count", len(tournaments)))

	dtos := models.TournamentDTOs(tournaments)
	err = response.JSONResponse(w, http.StatusOK, dtos)
	if err != nil {
		log.Error("failed to send tournaments response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}

func GetTournamentByIDHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetTournamentByIDHandler]")
	log.Debug("started")

	// /api/tournaments/{id}
	tournamentID, err := utils.GetInt64PathParameter(r.URL, 3)
	if err != nil {
		log.Warn("failed to get tournament id", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	log = log.With(zap.Int64("TournamentID", tournamentID))

	player, err := services.TournamentByIDGet(db.DB, tournamentID)
	if err != nil {
		log.Warn("failed to get tournament by id", zap.Error(err))
		// FIXME: parse the error and return 500 if something is wrong with database
		response.NewNotFound(w)
		return
	}

	err = response.JSONResponse(w, http.StatusOK, player.ToDTO())
	if err != nil {
		log.Error("failed to send tournament response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
