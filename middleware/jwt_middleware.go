package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/config"
	"neeft_back/app/models/users"
	"neeft_back/database"
	"neeft_back/utils"
)

func FindUserByClaim(claims config.JWTClaims, user *users.User) error {
	database.Database.Db.Find(&user, "id = ?", claims.UserId)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func VerifyJWT(c *fiber.Ctx) error {
	claims := config.JWTClaims{}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user := users.User{}

	if err := FindUserByClaim(claims, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return nil
}
