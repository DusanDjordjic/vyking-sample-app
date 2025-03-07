package tournaments_handlers

import (
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/models"
	"app/pkg/router/response"
	"app/pkg/services"
	"encoding/json"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type createTournamentPayload struct {
	Name      string          `json:"name"`
	Prize     float64         `json:"prize"`
	StartDate models.Datetime `json:"start_date"`
	EndDate   models.Datetime `json:"end_date"`
}

func (payload *createTournamentPayload) Parse(r *http.Request) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}

	// Make sure times are in UTC
	payload.StartDate = models.Datetime(time.Time(payload.StartDate).UTC())
	payload.EndDate = models.Datetime(time.Time(payload.EndDate).UTC())

	return nil
}

func (payload createTournamentPayload) Validate(r *http.Request) []string {
	var errs []string

	if len(payload.Name) == 0 {
		errs = append(errs, "name is required")
	}

	if payload.Prize <= 0 {
		errs = append(errs, "prize must be positive")
	}

	if time.Time(payload.StartDate).Compare(time.Now()) <= 0 {
		errs = append(errs, "start date can't be in past")
	}

	if time.Time(payload.EndDate).Compare(time.Time(payload.StartDate)) <= 0 {
		errs = append(errs, "end date must be after start date")
	}

	return errs
}

func CreateTournamentHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[CreateTournamentHandler]")
	log.Debug("started")

	payload := createTournamentPayload{}
	err := payload.Parse(r)
	if err != nil {
		log.Warn("failed to parse payload", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	errs := payload.Validate(r)
	if len(errs) > 0 {
		log.Warn("payload is not valid", zap.Strings("Errros", errs))
		response.NewBadRequest(w, errs...)
		return
	}

	data := services.TournamentCreateData{
		Name:      payload.Name,
		Prize:     payload.Prize,
		StartDate: time.Time(payload.StartDate),
		EndDate:   time.Time(payload.EndDate),
	}

	tournament, err := services.TournamentCreate(db.DB, data)
	if err != nil {
		log.Error("failed to create tournament", zap.Error(err))
		response.NewInternalError(w)
		return
	}

	err = response.JSONResponse(w, http.StatusCreated, tournament.ToDTO())
	if err != nil {
		log.Error("failed to send tournament response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
