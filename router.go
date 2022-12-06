package main

import (
	"github.com/gin-gonic/gin"
	"neeft_back/controllers"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Accueil)
	r.POST("/connect", controllers.Connect)
	r.POST("/newTeam", controllers.NewTeam)
	r.POST("/newTournament", controllers.NewTournament)
	r.POST("/register", controllers.Register)
	r.POST("/editTournament", controllers.EditTournament)

	// CORS OPTIONS requests
	r.OPTIONS("/connect", controllers.ConnectOptions)
	r.OPTIONS("/newTeam", controllers.NewTeamOptions)
	r.OPTIONS("/newTournament", controllers.NewTournament)
	r.OPTIONS("/register", controllers.RegisterOptions)
	r.OPTIONS("/editTournament", controllers.EditTournamentOptions)
}
