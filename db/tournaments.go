package db

import (
	"context"
	"neeft_back/models"
	"time"
)

func FetchTournamentById(id int) (*models.Tournament, error) {
	db := OpenDB()

	row := db.QueryRow("select * from tournaments where id=?", id)

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

func UpdateTournament(id int, name string, count int, price int, game string, teamsCount int, mode string, beginDate string, beginTime string) error {
	originalTournament, err := FetchTournamentById(id)
	if err != nil {
		return err
	}

	if len(name) == 0 {
		name = originalTournament.Name
	}

	if count == 0 {
		count = originalTournament.Count
	}

	if price == 0 {
		price = originalTournament.Price
	}

	if len(game) == 0 {
		game = originalTournament.Game
	}

	if teamsCount == 0 {
		teamsCount = originalTournament.TeamsCount
	}

	if len(mode) == 0 {
		mode = originalTournament.Mode
	}

	if len(beginDate) == 0 {
		beginDate = originalTournament.BeginDate
	}

	if len(beginTime) == 0 {
		beginTime = originalTournament.BeginTime
	}

	db := OpenDB()

	query := "UPDATE tournaments SET name = ?, count = ?, price = ?, game = ?, nbr_teams = ?, mode = ?, begin_date = ?, begin_time = ? WHERE id = ?"

	ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		name,
		count,
		price,
		game,
		teamsCount,
		mode,
		beginDate,
		beginTime,
		id)

	db.Close()
	cancelFunction()
	stmt.Close()

	return err
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
