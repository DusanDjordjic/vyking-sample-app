package services

import (
	"app/pkg/db/queries"
	"app/pkg/models"
	"database/sql"
	"errors"
	"fmt"
)

func PlayersGet(db *sql.DB, limit int64, offset int64) ([]models.Player, error) {
	var players []models.Player

	rows, err := db.Query(queries.PlayerGetAll, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get players, %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		v := models.Player{}

		err = rows.Scan(&v.ID, &v.Name, &v.Email, &v.AccountBalance)
		if err != nil {
			return nil, fmt.Errorf("failed to scan player, %s", err)
		}

		players = append(players, v)
	}

	return players, nil
}

// FIXME: return custom error so we can differentiate when user doesn't exists or something is wrong with database
func PlayersGetByID(db *sql.DB, id int64) (models.Player, error) {
	var player models.Player

	row := db.QueryRow(queries.PlayerGetByID, id)
	err := row.Scan(&player.ID, &player.Name, &player.Email, &player.AccountBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return player, fmt.Errorf("player doesn't exist")
		}

		return player, fmt.Errorf("failed to get player, %s", err)
	}
	return player, nil
}

type PlayerCreateData struct {
	Name           string
	Email          string
	AccountBalance float64
}

func PlayerCreate(db *sql.DB, data PlayerCreateData) (models.Player, error) {
	var player models.Player

	player.Name = data.Name
	player.Email = data.Email
	player.AccountBalance = data.AccountBalance

	res, err := db.Exec(queries.PlayerInsert, player.Name, player.Email, player.AccountBalance)
	if err != nil {
		return player, fmt.Errorf("failed to insert player, %s", err)
	}

	newID, err := res.LastInsertId()
	if err != nil {
		return player, fmt.Errorf("failed to get new player's id, %s", err)
	}

	player.ID = newID

	return player, nil
}

type PlayerUpdateData struct {
	PlayerID       int64
	AccountBalance float64
}

func PlayerUpdate(db *sql.DB, data PlayerUpdateData) error {
	res, err := db.Exec(queries.PlayerUpdateBalance, data.AccountBalance, data.PlayerID)
	if err != nil {
		return fmt.Errorf("failed to insert player, %s", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected, %s", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("plyare with id %d doesn't exist", data.PlayerID)
	}

	return nil
}
