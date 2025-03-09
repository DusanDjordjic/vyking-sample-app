package tournaments_handlers

import (
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/router/response"
	"app/pkg/services"
	"app/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type betOnTournamentPayload struct {
	PlayerID int64   `json:"player_id"`
	Amount   float64 `json:"amount"`
}

func (payload *betOnTournamentPayload) Parse(r *http.Request) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}

	return nil
}

func (payload betOnTournamentPayload) Validate(r *http.Request) []string {
	var errs []string

	if payload.Amount <= 0 {
		errs = append(errs, "amount must be positive")
	}

	return errs
}

func BetOnTournamentHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[GetTournamentByIDHandler]")
	log.Debug("started")

	// /api/tournaments/{id}/bet
	tournamentID, err := utils.GetInt64PathParameter(r.URL, 3)
	if err != nil {
		log.Warn("failed to get tournament id", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	log = log.With(zap.Int64("TournamentID", tournamentID))

	payload := betOnTournamentPayload{}
	err = payload.Parse(r)
	if err != nil {
		log.Warn("failed to parse payload", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	// TODO: when I add sessions make sure to get player id from that and not payload
	log = log.With(zap.Int64("PlayerID", payload.PlayerID))

	errs := payload.Validate(r)
	if len(errs) > 0 {
		log.Warn("payload is not valid", zap.Strings("Errros", errs))
		response.NewBadRequest(w, errs...)
		return
	}

	data := services.TournamentBetOnData{
		PlayerID:     payload.PlayerID,
		TournamentID: tournamentID,
		Amount:       payload.Amount,
	}

	newBet, err := services.TournamentBetOn(db.DB, data)
	if err != nil {
		log.Warn("failed to create a bet", zap.Error(err))
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			sqlState := string(mysqlErr.SQLState[:])
			switch sqlState {
			case "45001":
				response.NewBadRequest(w, "insufficient funds")
				return
			case "45002":
				response.NewBadRequest(w, "tournament doesn't exist")
				return
			case "45003":
				response.NewBadRequest(w, "tournament hasn't started yet")
				return
			case "45004":
				response.NewBadRequest(w, "tournament has already ended")
				return
			}
		}

		// TODO: tournament id may be invalid, handle that error
		// TODO: tournament isn't started yet or it has already ended

		// Don't know what went wrong
		response.NewInternalError(w)
		return
	}

	err = response.JSONResponse(w, http.StatusOK, newBet.ToDTO())
	if err != nil {
		log.Error("failed to send tournament response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
