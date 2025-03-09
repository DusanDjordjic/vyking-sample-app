package models

import "time"

type TournamentBet struct {
	ID           int64
	CreatedAt    time.Time
	PlayerID     int64
	TournamentID int64
	Amount       float64
}

type TournamentBetDTO struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	PlayerID     int64     `json:"player_id"`
	TournamentID int64     `json:"tournament_id"`
	Amount       float64   `json:"amount"`
}

func (v TournamentBet) ToDTO() TournamentBetDTO {
	return TournamentBetDTO{
		ID:           v.ID,
		CreatedAt:    v.CreatedAt,
		PlayerID:     v.PlayerID,
		TournamentID: v.TournamentID,
		Amount:       v.Amount,
	}
}

func TournamentBetDTOs(vs []TournamentBet) []TournamentBetDTO {
	dtos := make([]TournamentBetDTO, 0, len(vs))
	for _, v := range vs {
		dtos = append(dtos, v.ToDTO())
	}
	return dtos
}
