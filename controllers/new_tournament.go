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

func NewTournament(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")

	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")

	tournamentName := c.PostForm("name")
	tournamentGame := c.PostForm("game")
	tournamentPrice := c.PostForm("price")
	tournamentTeamCount := c.PostForm("nbr_teams")
	teamCount, _ := strconv.Atoi(tournamentTeamCount)
	parsedPrice, err := strconv.Atoi(tournamentPrice)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid price provided", "code": 401})
		db.Close()
		return
	}

	// Check if the tournament already exists
	row := db.QueryRow("select * from tournaments where name=? and game=?", tournamentName, tournamentGame)
	tournament := new(models.Tournament)
	err = row.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game, &tournament.TeamsCount, &tournament.IsFinished, &tournament.Mode)
	if err == nil && tournament.Name == tournamentName && tournament.IsFinished == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "A tournament with the same name already exists and isn't finished", "code": 401})
		db.Close()
		return
	}

	if tournamentGame == "Lol" {
		if teamCount%2 == 0 {
			// Insert an element in a table
			query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode) VALUES (?, ?, ?, ?, ?, ?, ?)"
			ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFunction()
			stmt, err := db.PrepareContext(ctx, query)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
				db.Close()
				return
			}
			defer stmt.Close()

			_, err = stmt.ExecContext(ctx, tournamentName, 1, parsedPrice, tournamentGame, 0, false, "unsupported")
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": "Failed", "error": err.Error(), "code": 401})
				db.Close()
				return
			}
		}
	} else if tournamentGame == "Fortnite" {
		tournamentMode := c.PostForm("mode")

		if tournamentMode != "solo" && tournamentMode != "duo" && tournamentMode != "trio" && tournamentMode != "squad" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid party mode", "code": 401})
			db.Close()
			return
		}

		// Insert an element in a table
		query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode) VALUES (?, ?, ?, ?, ?, ?, ?)"
		ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunction()
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
			db.Close()
			return
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, tournamentName, 1, parsedPrice, tournamentGame, 0, false, tournamentMode)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Failed", "error": err.Error(), "code": 401})
			db.Close()
			return
		}
	}

	// Get the tournament's infos
	getIdRow := db.QueryRow("select * from tournaments where name=? and game=? order by id desc", tournamentName, tournamentGame)
	tournament = new(models.Tournament)
	err = getIdRow.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game, &tournament.TeamsCount, &tournament.IsFinished, &tournament.Mode)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error(), "code": 401})
		db.Close()
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"message": "Success", "tournament_id": tournament.Id, "code": 200})

	db.Close()
}
