package models

type Player struct {
	ID             int64
	Name           string
	Email          string
	AccountBalance float64
}

type PlayerDTO struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	AccountBalance float64 `json:"account_balance"`
}

func (v Player) ToDTO() PlayerDTO {
	return PlayerDTO{
		ID:             v.ID,
		Name:           v.Name,
		Email:          v.Email,
		AccountBalance: v.AccountBalance,
	}
}

func PlayerDTOs(vs []Player) []PlayerDTO {
	dtos := make([]PlayerDTO, 0, len(vs))
	for _, v := range vs {
		dtos = append(dtos, v.ToDTO())
	}
	return dtos
}
