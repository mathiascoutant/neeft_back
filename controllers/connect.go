package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/db"
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

	registerUser, err := db.FetchUser(inUsername)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(registerUser.Password), []byte(inPassword)) != nil {
		utils.SendError(c, 401, utils.UsernameOrPasswordInvalidError)
		return
	} else {
		utils.SendOK(c, gin.H{"message": "Success", "userId": registerUser.Id})
	}
}
