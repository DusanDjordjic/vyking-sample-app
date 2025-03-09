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

type updatePlayerAccountBalancePayload struct {
	AccountBalance float64 `json:"account_balance"`
}

func (payload *updatePlayerAccountBalancePayload) Parse(r *http.Request) error {
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}

	return nil
}

func (payload updatePlayerAccountBalancePayload) Validate() []string {
	errs := make([]string, 0)

	if payload.AccountBalance < 0 {
		errs = append(errs, "account balance cannot be negative")
	}

	return errs
}

type updatePlayerBalanceResponse struct {
	NewAccountBalance float64 `json:"new_account_balance"`
}

func UpdatePlayerAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Log.Named("[CreatePlayerHandler]")
	log.Debug("started")

	// /api/players/{id}
	playerID, err := utils.GetInt64PathParameter(r.URL, 3)
	if err != nil {
		log.Warn("failed to get player id", zap.Error(err))
		response.NewBadRequest(w)
		return
	}

	log = log.With(zap.Int64("PlayerID", playerID))

	payload := updatePlayerAccountBalancePayload{}
	err = payload.Parse(r)
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

	data := services.PlayerUpdateData{
		PlayerID:       playerID,
		AccountBalance: payload.AccountBalance,
	}

	err = services.PlayerUpdate(db.DB, data)
	if err != nil {
		log.Error("failed to update balance player", zap.Error(err))
		// TODO: check the error and modify the response based on it
		response.NewBadRequest(w)
		return
	}

	res := updatePlayerBalanceResponse{NewAccountBalance: data.AccountBalance}

	err = response.JSONResponse(w, http.StatusCreated, res)
	if err != nil {
		log.Error("failed to send new balance response", zap.Error(err))
		// Status code is already set in response.JSONResponse
		return
	}

	log.Debug("finished")
}
