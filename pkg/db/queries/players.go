package queries

import (
	_ "embed"
)

//go:embed players/get_all.sql
var PlayerGetAll string

//go:embed players/get_by_id.sql
var PlayerGetByID string

//go:embed players/insert.sql
var PlayerInsert string

//go:embed players/update_balance.sql
var PlayerUpdateBalance string
