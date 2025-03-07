package services

import (
	"app/pkg/models"
	"database/sql"
	"errors"
	"fmt"
)

func PlayersGet(db *sql.DB, limit int64, offset int64) ([]models.Player, error) {
	var players []models.Player

	rows, err := db.Query("SELECT id, name, email, account_balance FROM players LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to select players, %s", err)
	}

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

	row := db.QueryRow("SELECT id, name, email, account_balance FROM players WHERE id = ?", id)
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

	res, err := db.Exec("INSERT INTO players (name, email, account_balance) VALUES (?, ?, ?)", player.Name, player.Email, player.AccountBalance)
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
