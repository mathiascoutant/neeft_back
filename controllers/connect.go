package controllers

import (
	"database/sql"
	"neeft_back/models"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(c *gin.Context) {
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")

	askedUsername := c.PostForm("username")
	askedPassword := c.PostForm("password")

	if askedUsername != "" {
		if askedPassword != "" {
			row := db.QueryRow("select * from users where username=? and password=?", askedUsername, askedPassword)
			user := new(models.User)
			err := row.Scan(&user.Id, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.EmailVerifiedAt)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": "Username or password is invalid", "code": 401})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"username": askedUsername, "id": user.Id, "password": askedPassword, "code": 200})
			}
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "empty password", "code": 500})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"message": "empty Username", "code": 500})
	}

	db.Close()
}
