package helper

/**
 * @Author: ANYARONKE Dare Samuel
 */

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// HandleErr is a helper function to handle errors
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// HashAndSalt is a helper function to hash and salt a password
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	HandleErr(err)
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

func CheckEmail(email string) (bool, error) {
	if len(email) < 3 && len(email) > 254 {
		return false, errors.New("L'email est trop court ou trop long")
	}
	if m, _ := regexp.MatchString("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$", email); !m {
		return false, errors.New("L'email n'est pas valide")
	}
	return true, nil
}

//---------------------ERRORS---------------------//

func Return200(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func Return400(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

func Return401(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": message})
}

func Return403(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": message})
}

func Return404(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": message})
}

func Return500(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}

func Return501(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"message": message})
}

func Return503(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"message": message})
}

func Return504(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{"message": message})
}
