package services

import (
	"app/pkg/db/queries"
	"app/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func TournamentsGet(db *sql.DB, limit int64, offset int64) ([]models.Tournament, error) {
	var tournaments []models.Tournament
	rows, err := db.Query(queries.TournamentGetAll, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get tournaments, %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		v := models.Tournament{}

		err = rows.Scan(&v.ID, &v.Name, &v.Prize, &v.StartDate, &v.EndDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan tournament, %s", err)
		}

		tournaments = append(tournaments, v)
	}

	return tournaments, nil
}

func TournamentByIDGet(db *sql.DB, id int64) (models.Tournament, error) {
	var tournament models.Tournament

	row := db.QueryRow(queries.TournamentGetbyID, id)
	err := row.Scan(&tournament.ID, &tournament.Name, &tournament.Prize, &tournament.StartDate, &tournament.EndDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tournament, fmt.Errorf("tournament doesn't exist")
		}

		return tournament, fmt.Errorf("failed to get tournament, %s", err)
	}

	return tournament, nil
}

type TournamentCreateData struct {
	Name      string
	Prize     float64
	StartDate time.Time
	EndDate   time.Time
}

func TournamentCreate(db *sql.DB, data TournamentCreateData) (models.Tournament, error) {
	var tournament models.Tournament
	tournament.Name = data.Name
	tournament.Prize = data.Prize
	tournament.StartDate = data.StartDate
	tournament.EndDate = data.EndDate

	res, err := db.Exec(queries.TournamentInsert, tournament.Name, tournament.Prize, tournament.StartDate, tournament.EndDate)
	if err != nil {
		return tournament, fmt.Errorf("failed to insert tournament, %s", err)
	}

	newID, err := res.LastInsertId()
	if err != nil {
		return tournament, fmt.Errorf("failed to get new tournament's id, %s", err)
	}

	tournament.ID = newID

	return tournament, nil
}

type TournamentBetOnData struct {
	TournamentID int64
	PlayerID     int64
	Amount       float64
}

func TournamentBetOn(db *sql.DB, data TournamentBetOnData) (models.TournamentBet, error) {
	row := db.QueryRow("CALL BetOnTournament(?, ?, ?)", data.PlayerID, data.TournamentID, data.Amount)
	bet := models.TournamentBet{}

	err := row.Scan(&bet.ID, &bet.CreatedAt, &bet.PlayerID, &bet.TournamentID, &bet.Amount)
	if err != nil {
		return bet, err
	}

	return bet, nil
}
