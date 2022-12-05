package db

import (
	"context"
	"neeft_back/models"
	"time"
)

func FetchTournament(name, game string) (*models.Tournament, error) {
	db := OpenDB()

	// Check if the tournament already exists
	row := db.QueryRow("select * from tournaments where name=? and game=? order by id desc", name, game)
	tournament := new(models.Tournament)
	err := row.Scan(&tournament.Id,
		&tournament.Name,
		&tournament.Count,
		&tournament.Price,
		&tournament.Game,
		&tournament.TeamsCount,
		&tournament.IsFinished,
		&tournament.Mode,
		&tournament.BeginDate,
		&tournament.BeginTime)

	db.Close()
	return tournament, err
}

func RegisterTournament(tournament models.Tournament) error {
	db := OpenDB()

	query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode, begin_date, begin_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		tournament.Name,
		tournament.Count,
		tournament.Price,
		tournament.Game,
		tournament.TeamsCount,
		tournament.IsFinished,
		tournament.Mode,
		tournament.BeginDate,
		tournament.BeginTime)

	db.Close()
	cancelFunction()
	stmt.Close()

	return err
}
