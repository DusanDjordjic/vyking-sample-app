package tournaments_handlers

import (
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/models"
	"app/pkg/router/response"
	"app/pkg/services"
	"app/pkg/utils"
	"net/http"

	"go.uber.org/zap"
)

func GetRankingForSingleTournament(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetRankingForSingleTournament]")
	log.Debug("started")

	// /api/tournaments/{id}/bet
	tournamentID, err := utils.GetInt64PathParameter(r.URL, 3)
	if err != nil {
		log.Warn("failed to get tournament id", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	log = log.With(zap.Int64("TournamentID", tournamentID))

	rankings, err := services.TournamentRankings(db.DB, tournamentID)
	if err != nil {
		log.Error("failed to get tournament rankings", zap.Error(err))
		response.NewInternalError(w)
		return
	}

	err = response.JSONResponse(w, http.StatusOK, models.TournamentRankingDTOs(rankings))
	if err != nil {
		log.Error("failed to send tournament rankings response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}

func GetRankingForAllTournaments(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetRankingForAllTournaments]")
	log.Debug("started")

	rankings, err := services.TournamentAllRankings(db.DB)
	if err != nil {
		log.Error("failed to get tournament rankings", zap.Error(err))
		response.NewInternalError(w)
		return
	}

	err = response.JSONResponse(w, http.StatusOK, models.TournamentRankingDTOs(rankings))
	if err != nil {
		log.Error("failed to send tournament rankings response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
