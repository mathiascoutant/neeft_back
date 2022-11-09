package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"neeft_back/models"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewTeam(c *gin.Context) {
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")

	teamName := c.PostForm("name")
	teamGame := c.PostForm("game")
	teamCreator := c.PostForm("creator")

	if len(teamName) == 0 || len(teamGame) == 0 || len(teamCreator) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid informations provided", "code": 401})
		db.Close()
		return
	}

	// Check if the team already exists
	row := db.QueryRow("select * from teams where team_name=?", teamName)
	team := new(models.Team)
	err := row.Scan(&team.Id, &team.Name, &team.Count, &team.GameName, &team.NbrTournoi, &team.CreatorName, &team.CreationDate)
	if err == nil && team.Name == teamName {
		c.JSON(http.StatusForbidden, gin.H{"message": "A team with the same name already exists", "code": 401})
		db.Close()
		return
	}

	// Insert an element in a table
	query := "INSERT INTO teams(team_name, count, game, nbr_tournoi, name_creator, date_creator) VALUES (?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
		db.Close()
		return
	}
	defer stmt.Close()

	curTime := time.Now()

	_, err = stmt.ExecContext(ctx, teamName, 1, teamGame, 0, teamCreator, strconv.Itoa(curTime.Day())+"/"+strconv.Itoa(int(curTime.Month()))+"/"+strconv.Itoa(curTime.Year()))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Failed", "error": err.Error(), "code": 401})
		db.Close()
		return
	}

	c.JSON(200, gin.H{"message": "Success", "team_name": teamName, "team_game": teamGame, "code": 200})

	db.Close()
}
