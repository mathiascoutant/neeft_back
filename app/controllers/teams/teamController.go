package teams

/**
 * @Author: ANYARONKE Daré Samuel
 */

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	usersController "neeft_back/app/controllers/users"
	"neeft_back/app/models/badges"
	"neeft_back/app/models/teams"
	usersModel "neeft_back/app/models/users"
	"neeft_back/database"
	"time"
)

// TeamSerialize  User : this is the router for the users not the model of User
// TeamSerialize serializer
type TeamSerialize struct {
	ID              uint            `json:"id"`
	User            usersModel.User `json:"user"`
	Name            string          `json:"name"`
	UserCount       uint            `json:"userCount"`
	GameName        string          `json:"gameName"`
	TournamentCount uint            `json:"tournamentCount"`
}

// CreateResponseTeam  /**
func CreateResponseTeam(userModel usersModel.User, teamModel teams.Team) TeamSerialize {
	return TeamSerialize{
		ID:              teamModel.ID,
		User:            userModel,
		Name:            teamModel.Name,
		UserCount:       teamModel.UserCount,
		GameName:        teamModel.GameName,
		TournamentCount: teamModel.TournamentCount,
	}
}

// CreateTeam function to create a team
func CreateTeam(c *fiber.Ctx) error {
	var team teams.Team

	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user usersModel.User
	if err := usersController.FindUser(team.UserId, &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	database.Database.Db.Create(&team)

	responseTeam := CreateResponseTeam(user, team)
	return c.Status(200).JSON(responseTeam)
}

// GetAllTeam function to get all teams
func GetAllTeam(c *fiber.Ctx) error {
	var teamsModel []teams.Team
	database.Database.Db.Find(&teamsModel)
	var responseTeams []TeamSerialize
	for _, team := range teamsModel {
		var user usersModel.User
		database.Database.Db.Find(&user, "id = ?", team.UserId)
		responseTeam := CreateResponseTeam(user, team)
		responseTeams = append(responseTeams, responseTeam)
	}

	return c.Status(200).JSON(responseTeams)
}

// FindTeam function to update a team
func FindTeam(id int, team *teams.Team) error {
	database.Database.Db.Find(&team, "id = ?", id)
	if team.ID == 0 {
		return errors.New("team does not exist")
	}
	return nil
}

// GetTeam function to get a team
func GetTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var team teams.Team

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindTeam(id, &team); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user usersModel.User
	database.Database.Db.First(&user, team.UserId)
	responseTeam := CreateResponseTeam(user, team)

	return c.Status(200).JSON(responseTeam)
}

// DeleteTeam function to delete a team
func DeleteTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var team teams.Team

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindTeam(id, &team); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Delete(&team)

	return c.Status(200).JSON("Team deleted successfully")
}

func AddBadgeTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var team teams.Team

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindTeam(id, &team); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var result badges.BadgeInput
	var badgeTemp badges.Badge

	if err := c.BodyParser(&result); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// Check if the user already has this badge
	if err := database.Database.Db.Where("recipient = ? AND recipient_id = ? AND title = ?", badges.RecipientTeam, id, result.Title).First(&badgeTemp).Error; err == nil {
		return c.Status(400).JSON("This team already have this badge")
	}

	if len(result.Title) == 0 || len(result.Description) == 0 {
		return c.Status(400).JSON("Invalid name or description")
	}

	badge := badges.Badge{
		Recipient:   badges.RecipientTeam,
		RecipientId: uint(id),
		Title:       result.Title,
		Description: result.Description,
		Category:    result.Category,
		Section:     result.Section,
		ReceiveDate: time.Now(),
	}

	type BadgeResponse struct {
		Id          uint      `json:"id"`
		Title       string    `json:"title"`
		Section     uint      `json:"section"`
		Category    uint      `json:"category"`
		Image       string    `json:"image"`
		ReceiveDate time.Time `json:"receive_date"`
	}

	resp := BadgeResponse{
		Id:          badge.ID,
		Title:       badge.Title,
		Section:     badge.Section,
		Category:    badge.Category,
		Image:       badge.Image,
		ReceiveDate: badge.ReceiveDate,
	}

	database.Database.Db.Create(&badge)
	return c.Status(200).JSON(resp)
}
