package controllers

import (
	"github.com/gin-gonic/gin"
	db2 "neeft_back/db"
	"neeft_back/models"
	"neeft_back/utils"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CreateTeamDTO struct {
	Name    string `json:"name"`
	Game    string `json:"game"`
	Creator string `json:"creator"`
}

func NewTeam(c *gin.Context) {
	utils.NewTeamOptions(c)

	var req CreateTeamDTO

	if err := c.BindJSON(&req); err != nil {
		utils.SendError(c, utils.InvalidRequestFormat)
		return
	}

	teamName := req.Name
	teamGame := req.Game
	teamCreator := req.Creator

	if len(teamName) == 0 || len(teamGame) == 0 || len(teamCreator) == 0 {
		utils.SendError(c, utils.InvalidInfosProvidedError)
		return
	}

	// Check if the team already exists

	team, err := db2.FetchTeam(teamName)
	if err == nil && team.Name == teamName {
		utils.SendError(c, utils.TeamWithSameNameExistsError)
		return
	}

	curTime := time.Now()
	dateString := strconv.Itoa(curTime.Day()) + "/" + strconv.Itoa(int(curTime.Month())) + "/" + strconv.Itoa(curTime.Year())

	err = db2.RegisterTeam(models.Team{
		Name:            teamName,
		UserCount:       1,
		GameName:        teamGame,
		TournamentCount: 0,
		CreatorName:     teamCreator,
		CreationDate:    dateString,
	})

	if err != nil {
		utils.SendError(c, utils.InternalError)
		return
	}

	utils.SendOK(c, gin.H{"message": "Success", "team": team})
}
