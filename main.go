package main

import (
	"github.com/gin-gonic/gin"
	"neeft_back/controllers"
)

func main() {
	controllers.InitDatabase()

	r := gin.Default()

	r.GET("/", controllers.Accueil)
	r.POST("/connect", controllers.Connect)
	r.POST("/new_team", controllers.NewTeam)
	r.POST("/new_tournament", controllers.NewTournament)

	// CORS OPTIONS requests
	r.OPTIONS("/connect", controllers.ConnectOptions)
	r.OPTIONS("/new_team", controllers.NewTeamOptions)
	r.OPTIONS("/new_tournament", controllers.NewTournament)

	r.Run()
}
