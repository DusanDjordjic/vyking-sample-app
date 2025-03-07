package models

import "time"

type Tournament struct {
	ID        int64
	Name      string
	Prize     float64
	StartDate time.Time
	EndDate   time.Time
}

type TournamentDTO struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Prize     float64  `json:"prize"`
	StartDate Datetime `json:"start_date"`
	EndDate   Datetime `json:"end_date"`
}

func (v Tournament) ToDTO() TournamentDTO {
	return TournamentDTO{
		ID:        v.ID,
		Name:      v.Name,
		Prize:     v.Prize,
		StartDate: Datetime(v.StartDate),
		EndDate:   Datetime(v.EndDate),
	}
}

func TournamentDTOs(vs []Tournament) []TournamentDTO {
	dtos := make([]TournamentDTO, 0, len(vs))
	for _, v := range vs {
		dtos = append(dtos, v.ToDTO())
	}
	return dtos
}
