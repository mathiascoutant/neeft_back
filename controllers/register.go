package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/models"
	"net/http"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)


type CreateUserDTO struct {
	Username string `json:"username"`
	Password  string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email  string `json:"email"`
}

func Register(c *gin.Context) {
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")
	defer db.Close()

	// TODO: Accept in input username, password, first_name, last_name, email, email_verified_at
	// Also use bcrypt
	var req CreateUserDTO

    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusForbidden, gin.H{"message": "Expected JSON format", "code": 403})
        return
    }

	inUsername := req.Username
	inPassword := req.Password
	inFirstName := req.FirstName
	inLastName := req.LastName
	inEmail := req.Email

	if len(inUsername) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid username", "code": 401})
		return
	}

	if len(inPassword) < 4 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid password or password length too short (must be more than 4 characters long)", "code": 401})
		return
	}

	if len(inFirstName) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid first name", "code": 401})
		return
	}

	if len(inLastName) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid last name", "code": 401})
		return
	}

	if len(inEmail) == 0 || !strings.Contains(inEmail, "@") {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid email address", "code": 401})
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
		c.JSON(http.StatusForbidden, gin.H{"message": "An account with the same username or email has already been created", "code": 401})
		return
	}

	// Insert the user into the table
	query := "INSERT INTO users(username, password, first_name, last_name, email, email_verified_at) VALUES (?, ?, ?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
		db.Close()
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
		c.JSON(http.StatusForbidden, gin.H{"message": "Unknown error occurred", "error": err.Error(), "code": 401})
		db.Close()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "code": 200})
}
