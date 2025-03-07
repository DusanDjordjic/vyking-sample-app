package queries

import (
	_ "embed"
)

//go:embed tournaments/get_all.sql
var TournamentGetAll string

//go:embed tournaments/get_by_id.sql
var TournamentGetbyID string

//go:embed tournaments/insert.sql
var TournamentInsert string
