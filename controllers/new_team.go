package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"neeft_back/models"
	"neeft_back/utils"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewTeamOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type CreateTeamDTO struct {
	Name    string `json:"name"`
	Game    string `json:"game"`
	Creator string `json:"creator"`
}

func NewTeam(c *gin.Context) {
	NewTeamOptions(c)

	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	var req CreateTeamDTO

	if err := c.BindJSON(&req); err != nil {
		utils.SendError(c, 401, utils.InvalidRequestFormat)
		return
	}

	teamName := req.Name
	teamGame := req.Game
	teamCreator := req.Creator

	if len(teamName) == 0 || len(teamGame) == 0 || len(teamCreator) == 0 {
		utils.SendError(c, 401, utils.InvalidInfosProvided)
		return
	}

	// Check if the team already exists
	row := db.QueryRow("select * from teams where team_name=?", teamName)
	team := new(models.Team)
	err := row.Scan(&team.Id, &team.Name, &team.Count, &team.GameName, &team.NbrTournoi, &team.CreatorName, &team.CreationDate)
	if err == nil && team.Name == teamName {
		utils.SendError(c, 401, utils.TeamWithSameNameExistsError)
		return
	}

	// Insert an element in a table
	query := "INSERT INTO teams(team_name, count, game, nbr_tournoi, name_creator, date_creator) VALUES (?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		utils.SendError(c, 401, utils.DatabaseError)
		return
	}
	defer stmt.Close()

	curTime := time.Now()

	_, err = stmt.ExecContext(ctx, teamName, 1, teamGame, 0, teamCreator, strconv.Itoa(curTime.Day())+"/"+strconv.Itoa(int(curTime.Month()))+"/"+strconv.Itoa(curTime.Year()))
	if err != nil {
		utils.SendError(c, 401, utils.DatabaseError)
		return
	}

	utils.SendOK(c, gin.H{"message": "Success", "team_name": teamName, "team_game": teamGame})
}
