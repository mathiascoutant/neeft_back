package users

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
)

/*
 * @Author ANYARONKE
 */

type AddFriend struct {
	ID     uint        `json:"id"`
	User   models.User `json:"user"`
	Friend models.User `json:"friend"`
}

func CreateResponseUserFriend(addFriend models.AddFriend, user models.User, friend models.User) AddFriend {
	return AddFriend{ID: addFriend.ID, User: user, Friend: friend}
}

func CreateUserFriend(c *fiber.Ctx) error {
	var addFriend models.AddFriend

	if err := c.BodyParser(&addFriend); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user1 models.User

	if err := FindUser(addFriend.UserId, &user1); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user2 models.User

	if err := FindUser(addFriend.FriendId, &user2); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&addFriend)
	responseUserFriend := CreateResponseUserFriend(addFriend, user1, user2)

	return c.Status(200).JSON(responseUserFriend)
}

func FindUserFriend(id int, userFriend *models.AddFriend) error {
	database.Database.Db.Find(&userFriend, "user_id = ?", id)
	if userFriend.ID == 0 {
		return errors.New("userFriend does not exist")
	}
	return nil
}

func GetUserFriends(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	var userFriends []models.AddFriend
	database.Database.Db.Find(&userFriends, "user_id = ? and is_friend = ?", id, 1)
	var responseUserFriends []AddFriend

	for _, userFriend := range userFriends {
		var user1 models.User
		var user2 models.User
		database.Database.Db.Find(&user1, "id = ?", userFriend.UserId)
		database.Database.Db.Find(&user2, "id = ?", userFriend.FriendId)
		responseUserFriend := CreateResponseUserFriend(userFriend, user1, user2)
		responseUserFriends = append(responseUserFriends, responseUserFriend)
	}

	return c.Status(200).JSON(responseUserFriends)
}

func GetUserFriend(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	var userFriend models.AddFriend

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindUserFriend(id, &userFriend); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user1 models.User
	var user2 models.User
	//fmt.Println(userFriend.OwnerId)
	database.Database.Db.Find(&user1, userFriend.UserId)
	database.Database.Db.Find(&user2, userFriend.FriendId)
	responseUserFriend := CreateResponseUserFriend(userFriend, user1, user2)

	return c.Status(200).JSON(responseUserFriend)
}
