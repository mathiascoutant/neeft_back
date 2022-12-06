package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"neeft_back/db"
	"neeft_back/utils"
)

func EditTournamentOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

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
	EditTournamentOptions(c)

	var req EditTournamentDTO

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, 401, utils.InvalidRequestFormat)
		return
	}

	tournament, err := db.FetchTournamentById(req.Id)
	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, 401, utils.TournamentDoesNotExistError)
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
		utils.SendError(c, 401, utils.InternalError)
		return
	}

	tournament, err = db.FetchTournamentById(req.Id)

	if err != nil {
		fmt.Println(err.Error())
		utils.SendError(c, 401, utils.InternalError)
		return
	}

	utils.SendOK(c, gin.H{"message": "Success", "tournament": tournament})
}
