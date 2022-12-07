package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	db2 "neeft_back/db"
	"neeft_back/models"
	"neeft_back/utils"
	"strings"
)

type CreateUserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func Register(c *gin.Context) {
	utils.RegisterOptions(c)
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	// Accept in input username, password, first_name, last_name, email, email_verified_at
	// Also use bcrypt
	var req CreateUserDTO

	if err := c.BindJSON(&req); err != nil {
		utils.SendError(c, utils.InvalidRequestFormat)
		return
	}

	inUsername := req.Username
	inPassword := req.Password
	inFirstName := req.FirstName
	inLastName := req.LastName
	inEmail := req.Email

	if len(inUsername) == 0 {
		utils.SendError(c, utils.UsernameEmptyError)
		return
	}

	if len(inPassword) < 4 {
		utils.SendError(c, utils.PasswordTooShortError)
		return
	}

	if len(inFirstName) == 0 {
		utils.SendError(c, utils.InvalidFirstNameError)
		return
	}

	if len(inLastName) == 0 {
		utils.SendError(c, utils.InvalidLastNameError)
		return
	}

	if len(inEmail) == 0 || !strings.Contains(inEmail, "@") {
		utils.SendError(c, utils.InvalidEmailError)
		return
	}

	// Check if a user exists with the passed username or email
	row := db.QueryRow("select * from users where username=? or email=?", inUsername, inEmail)
	user := new(models.User)

	err := row.Scan(&user.Id,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EmailVerificationDate)

	if err == nil || user.Username == inUsername || user.Email == inEmail {
		utils.SendError(c, utils.AccountAlreadyExistError)
		return
	}

	toRegisterUser := models.User{
		Username:              inUsername,
		Password:              inPassword,
		FirstName:             inFirstName,
		LastName:              inLastName,
		Email:                 inEmail,
		EmailVerificationDate: 0,
	}
	err = db2.RegisterUser(toRegisterUser)

	if err != nil {
		fmt.Println("RegisterUser: " + err.Error())
		utils.SendError(c, utils.InternalError)
		return
	}

	utils.SendOK(c, gin.H{"message": "success", "userId": toRegisterUser.Id})
}
