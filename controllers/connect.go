package controllers

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/models"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Request-Headers", "Content-Type")
}

func Connect(c *gin.Context) {

	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	inUsername := c.PostForm("username")
	inPassword := c.PostForm("password")

	if len(inUsername) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Please enter a valid username", "code": 403})
		return
	} else if len(inPassword) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Please enter a non-null password", "code": 403})
		return
	}

	row := db.QueryRow("select * from users where username=?", inUsername)

	user := new(models.User)
	hashedPassword := ""

	err := row.Scan(&user.Id,
		&user.Username,
		&hashedPassword,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EmailVerifiedAt)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inPassword)) != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Username or password is invalid", "code": 401})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"username": inUsername, "id": user.Id, "code": 200})
	}
}
