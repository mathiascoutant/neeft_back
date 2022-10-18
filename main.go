package main

import (
	"neeft_back/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDatabase()

	r := gin.Default()

	r.GET("/", controllers.Accueil)

	r.Run()
}
