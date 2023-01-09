package users

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/helper"
	"neeft_back/app/models"
	"neeft_back/database"
	"time"
)

// CreateResponseUser /**
func CreateResponseUser(userModel models.User) models.User {
	return userModel
}

func EmailExist(email string) bool {
	var user models.User
	if err := database.Database.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

// CreateUser function to create a new user
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if EmailExist(user.Email) {
		return c.Status(400).JSON("Email already exist")
	}

	hashUserPassword := helper.HashAndSalt([]byte(user.Password))
	user.Password = hashUserPassword

	// Default values
	user.IsBan = false
	user.IsSuperAdmin = false

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

// GetAllUser function to get all users in the database
func GetAllUser(c *fiber.Ctx) error {
	var allUsers []models.User
	database.Database.Db.Find(&allUsers)
	var responseUsers []models.User
	for _, user := range allUsers {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

// FindUser function to find a user by id in the database
func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

// GetUser function to find a user by id in the database like find function
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// UpdateUser function is used to update a user
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		Username   string `json:"username"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
		BirthDate  string `json:"birth_date"`
		Avatar     string `json:"avatar"`
		Updated_at time.Time
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.LastName = updateData.Username
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.Email = updateData.Email
	user.BirthDate = updateData.BirthDate
	user.Avatar = updateData.Avatar
	user.Updated_at = updateData.Updated_at

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

// DeleteUser function to delete a user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindUser(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}
