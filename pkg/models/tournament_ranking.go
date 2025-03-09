package models

type TournamentRanking struct {
	PlayerID int64
	Prize    float64
}

type TournamentRankingDTO struct {
	PlayerID int64   `json:"player_id"`
	Prize    float64 `json:"prize"`
}

func (v TournamentRanking) ToDTO() TournamentRankingDTO {
	return TournamentRankingDTO{
		PlayerID: v.PlayerID,
		Prize:    v.Prize,
	}
}

func TournamentRankingDTOs(vs []TournamentRanking) []TournamentRankingDTO {
	dtos := make([]TournamentRankingDTO, 0, len(vs))
	for _, v := range vs {
		dtos = append(dtos, v.ToDTO())
	}
	return dtos
}
