package players_handlers

import (
	"app/pkg/db"
	"app/pkg/logger"
	"app/pkg/router/response"
	"app/pkg/services"
	"app/pkg/utils"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type createPlayerPayload struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float64 `json:"account_balance"`
}

func (payload *createPlayerPayload) Parse(r *http.Request) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}

	return nil
}

func (payload createPlayerPayload) Validate() []string {
	errs := make([]string, 0)

	if len(payload.Name) == 0 {
		errs = append(errs, "name is required")
	}

	if payload.AccountBalance < 0 {
		errs = append(errs, "account balance cannot be negative")
	}

	if len(payload.Email) == 0 {
		errs = append(errs, "email is required")
	} else {
		if !utils.IsEmailValid(payload.Email) {
			errs = append(errs, "email is not valid")
		}
	}

	return errs
}

func CreatePlayerHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[CreatePlayerHandler]")
	log.Debug("started")

	payload := createPlayerPayload{}
	err := payload.Parse(r)
	if err != nil {
		log.Warn("failed to parse payload", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	errs := payload.Validate()
	if len(errs) != 0 {
		log.Warn("payload is not valid", zap.Strings("Errors", errs))
		response.ErrorResponse(w, http.StatusBadRequest, errs...)
		return
	}

	data := services.PlayerCreateData{
		Name:           payload.Name,
		Email:          payload.Email,
		AccountBalance: payload.AccountBalance,
	}

	player, err := services.PlayerCreate(db.DB, data)
	if err != nil {
		log.Error("failed to create player", zap.Error(err))
		response.NewInternalError(w)
		return
	}

	err = response.JSONResponse(w, http.StatusCreated, player.ToDTO())
	if err != nil {
		log.Error("failed to send player response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
