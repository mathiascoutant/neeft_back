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

func Connect(c *gin.Context) {
	utils.ConnectOptions(c)

	var req ConnectRequestBody

	if err := c.BindJSON(&req); err != nil {
		utils.SendError(c, utils.InvalidRequestFormat)
		return
	}

	inUsername := req.Username
	inPassword := req.Password

	if len(inUsername) == 0 {
		utils.SendError(c, utils.UsernameEmptyError)
		return
	} else if len(inPassword) == 0 {
		utils.SendError(c, utils.PasswordEmptyError)
		return
	}

	registerUser, err := db.FetchUser(inUsername)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(registerUser.Password), []byte(inPassword)) != nil {
		utils.SendError(c, utils.UsernameOrPasswordInvalidError)
		return
	} else {
		utils.SendOK(c, gin.H{"message": "Success", "userId": registerUser.Id})
	}
}
