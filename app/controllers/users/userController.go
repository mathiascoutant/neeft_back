package users

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/helper"
	"neeft_back/app/models/users"
	"neeft_back/database"
	"time"
)

// UserSerialize User : this is the router for the users not the model of User
// UserSerialize serializer
type UserSerialize struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// CreateResponseUser /**
func CreateResponseUser(userModel users.User) UserSerialize {
	return UserSerialize{
		ID:        userModel.ID,
		Username:  userModel.Username,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Password:  userModel.Password,
	}
}

func EmailExist(email string) bool {
	var user users.User
	if err := database.Database.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

// CreateUser function to create a new user
func CreateUser(c *fiber.Ctx) error {
	var user users.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if EmailExist(user.Email) {
		return c.Status(400).JSON("Email already exist")
	}

	hashUserPassword := helper.HashAndSalt([]byte(user.Password))
	user.Password = hashUserPassword

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

// GetAllUser function to get all users in the database
func GetAllUser(c *fiber.Ctx) error {
	var allUsers []users.User
	database.Database.Db.Find(&allUsers)
	var responseUsers []UserSerialize
	for _, user := range allUsers {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

// FindUser function to find a user by id in the database
func FindUser(id int, user *users.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

// GetUser function to find a user by id in the database like find function
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user users.User

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

	var user users.User

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
		Password   string `json:"password"`
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
	user.Password = updateData.Password
	user.Updated_at = updateData.Updated_at

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

// DeleteUser function to delete a user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user users.User

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
