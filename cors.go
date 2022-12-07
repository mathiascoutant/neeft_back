package main

import (
	"github.com/gin-gonic/gin"
	"neeft_back/utils"
)

func setupCORS(r *gin.Engine) {
	r.OPTIONS("/connect", utils.ConnectOptions)
	r.OPTIONS("/newTeam", utils.NewTeamOptions)
	r.OPTIONS("/newTournament", utils.NewTournamentOptions)
	r.OPTIONS("/register", utils.RegisterOptions)
	r.OPTIONS("/editTournament", utils.EditTournamentOptions)
}
