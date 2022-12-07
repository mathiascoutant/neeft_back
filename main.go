package main

import (
	"github.com/gin-gonic/gin"
	"neeft_back/controllers"
)

func main() {
	controllers.InitDatabase()

	r := gin.Default()

	setupRoutes(r)
	setupCORS(r)

	r.Run()
}
