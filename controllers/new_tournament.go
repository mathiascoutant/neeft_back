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
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")

	tournamentName := c.PostForm("name")
	tournamentGame := c.PostForm("game")
	tournamentPrice := c.PostForm("price")

	if len(tournamentName) == 0 || len(tournamentGame) == 0 || len(tournamentPrice) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid informations provided", "code": 401})
		db.Close()
		return
	}

	// Check if the tournament already exists
	row := db.QueryRow("select * from tournaments where name=?", tournamentName)
	tournament := new(models.Tournament)
	err := row.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game)
	if err == nil && tournament.Name == tournamentName {
		c.JSON(http.StatusForbidden, gin.H{"message": "A tournament with the same name already exists", "code": 401})
		db.Close()
		return
	}

	// Insert an element in a table
	query := "INSERT INTO tournaments(name, count, price, game) VALUES (?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
		db.Close()
		return
	}
	defer stmt.Close()

	parsedPrice, err := strconv.Atoi(tournamentPrice)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid price provided", "code": 401})
		db.Close()
		return
	}

	_, err = stmt.ExecContext(ctx, tournamentName, 1, parsedPrice, tournamentGame)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Failed", "error": err.Error(), "code": 401})
		db.Close()
		return
	}

	c.JSON(200, gin.H{"message": "Success", "tournament_name": tournamentName, "tournament_game": tournamentGame, "tournament_moula": tournamentPrice, "code": 200})

	db.Close()
}
