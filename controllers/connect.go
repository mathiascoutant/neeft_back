package controllers

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/models"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type ConnectRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ConnectOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func Connect(c *gin.Context) {
	ConnectOptions(c)
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	var req ConnectRequestBody

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Expected JSON format", "code": 403})
		return
	}

	inUsername := req.Username
	inPassword := req.Password

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
