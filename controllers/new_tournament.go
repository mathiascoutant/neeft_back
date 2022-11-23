package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"neeft_back/models"
	"neeft_back/utils"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewTournamentOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type CreateTournamentDTO struct {
	Name     string `json:"name"`
	Game     string `json:"game"`
	Mode     string `json:"mode"`
	Price    int    `json:"price"`
	NbrTeams int    `json:"nbr_teams"`
}

func NewTournament(c *gin.Context) {
	NewTournamentOptions(c)

	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	var req CreateTournamentDTO

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, 401, utils.InvalidRequestFormat)
		return
	}

	tournamentName := req.Name
	tournamentGame := req.Game
	tournamentPrice := req.Price
	tournamentTeamCount := req.NbrTeams

	// Check if the tournament name isn't empty
	if len(tournamentName) <= 0 {
		utils.SendError(c, 401, utils.TournamentNameEmptyError)
		return
	}

	// Check if the tournament price is higher than 0
	if tournamentPrice <= 0 {
		utils.SendError(c, 401, utils.InvalidPriceError)
		return
	}

	// Check if the number of team is equals or higher than 2
	if tournamentTeamCount < 2 {
		utils.SendError(c, 401, utils.AtLeastTwoTeamsError)
		return
	}

	// Check if the tournament already exists
	row := db.QueryRow("select * from tournaments where name=? and game=?", tournamentName, tournamentGame)
	tournament := new(models.Tournament)
	err := row.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game, &tournament.TeamsCount, &tournament.IsFinished, &tournament.Mode)
	if err == nil && tournament.Name == tournamentName && tournament.IsFinished == 0 {
		utils.SendError(c, 401, utils.TournamentWithSameNameUnfinishedError)
		return
	}

	if tournamentGame == "Lol" {
		if tournamentTeamCount%2 != 0 {
			utils.SendError(c, 401, utils.InvalidTeamSize)
			return
		} else {
			// Insert an element in a table
			query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode) VALUES (?, ?, ?, ?, ?, ?, ?)"
			ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFunction()
			stmt, err := db.PrepareContext(ctx, query)
			if err != nil {
				utils.SendError(c, 401, utils.DatabaseError)
				return
			}
			defer stmt.Close()

			_, err = stmt.ExecContext(ctx, tournamentName, 1, tournamentPrice, tournamentGame, 0, false, "unsupported")
			if err != nil {
				utils.SendError(c, 401, utils.DatabaseError)
				return
			}
		}
	} else if tournamentGame == "Fortnite" {
		tournamentMode := req.Mode

		if tournamentMode != "solo" && tournamentMode != "duo" && tournamentMode != "trio" && tournamentMode != "squad" {
			utils.SendError(c, 401, utils.InvalidPartyMode)
			return
		}

		// Insert an element in a table
		query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode) VALUES (?, ?, ?, ?, ?, ?, ?)"
		ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunction()
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			utils.SendError(c, 401, utils.DatabaseError)
			return
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, tournamentName, 1, tournamentPrice, tournamentGame, 0, false, tournamentMode)
		if err != nil {
			utils.SendError(c, 401, utils.DatabaseError)
			return
		}
	}

	// Get the tournament's infos
	getIdRow := db.QueryRow("select * from tournaments where name=? and game=? order by id desc", tournamentName, tournamentGame)
	tournament = new(models.Tournament)
	err = getIdRow.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game, &tournament.TeamsCount, &tournament.IsFinished, &tournament.Mode)
	if err != nil {
		utils.SendError(c, 401, utils.DatabaseError)
		return
	}

	utils.SendOK(c, gin.H{"message": "Success", "tournament_id": tournament.Id})
}
