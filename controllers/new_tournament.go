package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"neeft_back/db"
	"neeft_back/models"
	"neeft_back/utils"
)

type CreateTournamentDTO struct {
	Name      string `json:"name"`
	Game      string `json:"game"`
	Mode      string `json:"mode"`
	Price     int    `json:"price"`
	TeamCount int    `json:"teamCount"`
	BeginTime string `json:"beginTime"`
	BeginDate string `json:"beginDate"`
}

func NewTournament(c *gin.Context) {
	utils.NewTournamentOptions(c)

	var req CreateTournamentDTO

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, utils.InvalidRequestFormat)
		return
	}

	tournamentName := req.Name
	tournamentGame := req.Game
	tournamentPrice := req.Price
	tournamentTeamCount := req.TeamCount
	tournamentBeginTime := req.BeginTime
	tournamentBeginDate := req.BeginDate

	// Check if the tournament name isn't empty
	if len(tournamentName) <= 0 {
		utils.SendError(c, utils.TournamentNameEmptyError)
		return
	}

	// Check if the tournament price is higher than 0
	if tournamentPrice <= 0 {
		utils.SendError(c, utils.InvalidPriceError)
		return
	}

	// Check if the number of team is equals or higher than 2
	if tournamentTeamCount < 2 {
		utils.SendError(c, utils.AtLeastTwoTeamsError)
		return
	}

	// Check if the begin & end time aren't null
	if len(tournamentBeginDate) <= 0 || len(tournamentBeginTime) <= 0 {
		utils.SendError(c, utils.InvalidDateTimeError)
		return
	}

	tournament, err := db.FetchTournament(tournamentName, tournamentGame)
	if err == nil && tournament.Name == tournamentName && tournament.IsFinished == 0 {
		utils.SendError(c, utils.TournamentWithSameNameUnfinishedError)
		return
	}

	tournamentMode := "none"

	if tournamentGame == "Lol" {
		if tournamentTeamCount%2 != 0 {
			utils.SendError(c, utils.InvalidTeamSizeError)
			return
		}
	} else if tournamentGame == "Fortnite" {
		tournamentMode = req.Mode

		if tournamentMode != "solo" && tournamentMode != "duo" && tournamentMode != "trio" && tournamentMode != "squad" {
			utils.SendError(c, utils.InvalidPartyModeError)
			return
		}
	}

	// Insert an element in the database
	err = db.RegisterTournament(models.Tournament{
		Name:       tournamentName,
		Count:      1,
		Price:      tournamentPrice,
		Game:       tournamentGame,
		TeamsCount: 0,
		IsFinished: 0,
		Mode:       tournamentMode,
		BeginDate:  tournamentBeginDate,
		BeginTime:  tournamentBeginTime,
	})

	if err != nil {
		fmt.Println("RegisterTournament: " + err.Error())
		utils.SendError(c, utils.InternalError)
		return
	}

	// Check if the tournament exists
	tournament, err = db.FetchTournament(tournamentName, tournamentGame)
	if err != nil {
		fmt.Println("FetchTournament: " + err.Error())
		utils.SendError(c, utils.InternalError)
		return
	}

	utils.SendOK(c, gin.H{"message": "Success", "tournament": tournament})
}
