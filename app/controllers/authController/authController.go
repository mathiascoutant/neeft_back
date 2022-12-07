package authController

/**
 * @Author ANYARONKE Dare Samuel
 */

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/helper"
	"neeft_back/app/models/users"
	"neeft_back/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"neeft_back/app/config"
)

// Login : Login a user and return a token to be used for authentication
func Login(c *fiber.Ctx) error {

	userInformation := new(users.User)

	// Get the user information from the request body
	if err := c.BodyParser(userInformation); err != nil {
		return helper.Return400(c, "Invalid user information")
	}

	// Check if the user exists in the database
	var user users.User
	if err := database.Database.Db.Find(&user, "email = ?", userInformation.Email).First(&user).Error; err != nil {
		return helper.Return401(c, "Invalid credentials")
	}

	// Check if the password is correct
	if err := helper.ComparePasswords(user.Password, []byte(userInformation.Password)); !err {
		return helper.Return401(c, "Invalid credentials")
	}

	// Generate JWT token for Auth user
	expireTime := time.Now().Add(time.Minute * 60)
	clams := &config.TWTClaim{
		Email:            user.Email,
		RegisteredClaims: jwt.RegisteredClaims{Issuer: "neeft", ExpiresAt: jwt.NewNumericDate(expireTime)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	tokenString, err := token.SignedString(config.JWT_SECRET)
	if err != nil {
		return helper.Return500(c, err.Error())
	}

	// Send token to cookie
	// send token to cookie
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expireTime,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login success",
		"token":   tokenString,
		"user":    user,
	})
}
