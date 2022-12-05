package db

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/models"
	"time"
)

func FetchUser(username string) (*models.User, error) {
	db := OpenDB()

	row := db.QueryRow("select * from users where username=?", username)

	user := new(models.User)

	err := row.Scan(&user.Id,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.EmailVerificationDate)

	db.Close()
	return user, err
}

func RegisterUser(user models.User) error {
	db := OpenDB()

	query := "INSERT INTO users(username, password, first_name, last_name, email, email_verified_at) VALUES (?, ?, ?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	_, err = stmt.ExecContext(ctx,
		user.Username,
		string(hashedPassword),
		user.FirstName,
		user.LastName,
		user.Email,
		0)

	db.Close()
	return err
}
