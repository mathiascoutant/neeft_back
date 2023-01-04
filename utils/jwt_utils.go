package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"neeft_back/app/config"
	"time"
)

// CheckJWT Checks if the passed JWT token is valid, and not expired
func CheckJWT(ctx *fiber.Ctx, decodedClaims *config.JWTClaims) error {
	token := ctx.Cookies("token", "")

	if token == "" {
		return errors.New("user not authenticated")
	}

	claims := config.JWTClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_SECRET, nil
	})

	if err != nil {
		return err
	}

	if tokenClaims, ok := jwtToken.Claims.(*config.JWTClaims); ok && jwtToken.Valid {
		*decodedClaims = *tokenClaims

		if time.Now().Unix() > decodedClaims.ExpiresAt.Unix() {
			return errors.New("token is expired")
		}
		return nil
	} else {
		return errors.New("invalid jwt token")
	}
}
