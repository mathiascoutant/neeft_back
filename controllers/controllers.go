package controllers

import (
	"neeft_back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Library []models.Accueil
var Counter int

func InitDatabase() {
	Counter = 1

	Accueil1 := models.Accueil{
		Id:        1,
		Titre:     "Ydays neeft",
		Createurs: "Kenan, Mathias",
	}
	Library = append(Library, Accueil1)

}

func Accueil(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
	c.JSON(http.StatusOK, gin.H{"data": Library})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input models.Accueil
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Counter++
	// Create book
	Accueil := models.Accueil{Id: Counter, Titre: input.Titre, Createurs: input.Createurs}
	Library = append(Library, Accueil)

	c.JSON(http.StatusOK, gin.H{"data": Accueil})
}
