package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/models"
	"neeft_back/utils"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type CreateUserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func Register(c *gin.Context) {
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	// Accept in input username, password, first_name, last_name, email, email_verified_at
	// Also use bcrypt
	var req CreateUserDTO

	if err := c.BindJSON(&req); err != nil {
		utils.SendError(c, 401, utils.InvalidRequestFormat)
		return
	}

	inUsername := req.Username
	inPassword := req.Password
	inFirstName := req.FirstName
	inLastName := req.LastName
	inEmail := req.Email

	if len(inUsername) == 0 {
		utils.SendError(c, 401, utils.UsernameEmptyError)
		return
	}

	if len(inPassword) < 4 {
		utils.SendError(c, 401, utils.PasswordTooShortError)
		return
	}

	if len(inFirstName) == 0 {
		utils.SendError(c, 401, utils.InvalidFirstNameError)
		return
	}

	if len(inLastName) == 0 {
		utils.SendError(c, 401, utils.InvalidLastNameError)
		return
	}

	if len(inEmail) == 0 || !strings.Contains(inEmail, "@") {
		utils.SendError(c, 401, utils.InvalidEmailError)
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
		&user.EmailVerifiedAt)

	if err == nil || user.Username == inUsername || user.Email == inEmail {
		utils.SendError(c, 401, utils.AccountAlreadyExistError)
		return
	}

	// Insert the user into the table
	query := "INSERT INTO users(username, password, first_name, last_name, email, email_verified_at) VALUES (?, ?, ?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		utils.SendError(c, 401, utils.DatabaseError)
		return
	}
	defer stmt.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inPassword), 10)
	_, err = stmt.ExecContext(ctx,
		inUsername,
		string(hashedPassword),
		inFirstName,
		inLastName,
		inEmail,
		0)

	if err != nil {
		utils.SendError(c, 401, utils.DatabaseError)
		return
	}

	utils.SendOK(c, gin.H{})
}
