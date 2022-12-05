package db

import (
	"context"
	"neeft_back/models"
	"time"
)

func FetchTeam(name string) (*models.Team, error) {
	db := OpenDB()

	row := db.QueryRow("select * from teams where team_name=?", name)
	team := new(models.Team)
	err := row.Scan(&team.Id,
		&team.Name,
		&team.UserCount,
		&team.GameName,
		&team.TournamentCount,
		&team.CreatorName,
		&team.CreationDate)

	db.Close()
	return team, err
}

func RegisterTeam(team models.Team) error {
	db := OpenDB()

	query := "INSERT INTO teams(team_name, count, game, nbr_tournoi, name_creator, date_creator) VALUES (?, ?, ?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		team.Name,
		team.UserCount,
		team.GameName,
		team.TournamentCount,
		team.CreatorName,
		team.CreationDate)

	db.Close()
	return err
}
