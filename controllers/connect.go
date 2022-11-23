package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/models"
	"neeft_back/utils"

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
		utils.SendError(c, 401, utils.InvalidRequestFormat)
		return
	}

	inUsername := req.Username
	inPassword := req.Password

	if len(inUsername) == 0 {
		utils.SendError(c, 401, utils.UsernameEmptyError)
		return
	} else if len(inPassword) == 0 {
		utils.SendError(c, 401, utils.PasswordEmptyError)
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
		utils.SendError(c, 401, utils.UsernameOrPasswordInvalidError)
		return
	} else {
		utils.SendOK(c, gin.H{"username": inUsername, "id": user.Id})
	}
}
