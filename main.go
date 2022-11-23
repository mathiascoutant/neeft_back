package main

import (
	"github.com/gin-gonic/gin"
	"neeft_back/controllers"
)

func main() {
	controllers.InitDatabase()

	r := gin.Default()

	// CORS part : we allow everyone
	r.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}(),
	)

	r.GET("/", controllers.Accueil)
	r.POST("/connect", controllers.Connect)
	r.POST("/new_team", controllers.NewTeam)
	r.POST("/new_tournament", controllers.NewTournament)

	r.Run()
}
