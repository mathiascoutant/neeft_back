package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"neeft_back/db"
	"neeft_back/utils"
)

type EditTournamentDTO struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Game      string `json:"game"`
	Mode      string `json:"mode"`
	Price     int    `json:"price"`
	TeamCount int    `json:"teamCount"`
	BeginTime string `json:"beginTime"`
	BeginDate string `json:"beginDate"`
}

func EditTournament(c *gin.Context) {
	utils.EditTournamentOptions(c)

	var req EditTournamentDTO

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, utils.InvalidRequestFormat)
		return
	}

	tournament, err := db.FetchTournamentByID(req.Id)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, utils.TournamentDoesNotExistError)
		return
	}

	tournamentId := req.Id
	tournamentName := req.Name
	tournamentGame := req.Game
	tournamentPrice := req.Price
	tournamentMode := req.Mode
	tournamentTeamCount := req.TeamCount
	tournamentBeginTime := req.BeginTime
	tournamentBeginDate := req.BeginDate

	err = db.UpdateTournament(tournamentId,
		tournamentName,
		0,
		tournamentPrice,
		tournamentGame,
		tournamentTeamCount,
		tournamentMode,
		tournamentBeginDate,
		tournamentBeginTime)

	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, utils.InternalError)
		return
	}

	tournament, err = db.FetchTournamentByID(req.Id)

	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, utils.InternalError)
		return
	}

	utils.SendOK(c, gin.H{"message": "Success", "tournament": tournament})
}
