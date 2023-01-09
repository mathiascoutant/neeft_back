package tournament

/**
 * @Author ANYARONKE
 */

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
)

// TournamentSerialize serializer
type TournamentSerialize struct {
	ID         uint   `json:"id" `
	Name       string `json:"name"`
	Count      uint   `json:"count"`
	Price      uint   `json:"price"`
	Game       string `json:"game"`
	TeamsCount uint   `json:"teamsCount"`
	IsFinished bool   `json:"isFinished"`
	Mode       string `json:"mode"`
}

// CreateResponseTournament /**
func CreateResponseTournament(tournamentModel models.Tournament) models.Tournament {
	return tournamentModel
}

// CreateTournament function to create a new tournament
func CreateTournament(c *fiber.Ctx) error {
	var tournament models.Tournament

	if err := c.BodyParser(&tournament); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Default values
	tournament.IsFinished = false

	database.Database.Db.Create(&tournament)
	responseTournament := CreateResponseTournament(tournament)
	return c.Status(200).JSON(responseTournament)
}

// GetAllTournament function to get all users in the database
func GetAllTournament(c *fiber.Ctx) error {
	var allTournament []models.Tournament
	database.Database.Db.Find(&allTournament)
	var responseTournaments []models.Tournament
	for _, tournament := range allTournament {
		responseTournament := CreateResponseTournament(tournament)
		responseTournaments = append(responseTournaments, responseTournament)
	}
	return c.Status(200).JSON(responseTournaments)
}

// FindTournament function to find a user by id in the database
func FindTournament(id int, tournament *models.Tournament) error {
	database.Database.Db.Find(&tournament, "id = ?", id)
	if tournament.ID == 0 {
		return errors.New("tournament does not exist")
	}
	return nil
}

// GetTournament function to find a user by id in the database like find function
func GetTournament(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tournament models.Tournament

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindTournament(id, &tournament); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseTournament(tournament)

	return c.Status(200).JSON(responseUser)
}

// UpdateTournament function is used to update a user
func UpdateTournament(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tournament models.Tournament

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindTournament(id, &tournament)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updateData models.Tournament

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// Resetting values that clients shouldn't modify
	updateData.IsFinished = tournament.IsFinished

	database.Database.Db.Save(&tournament)

	responseUser := CreateResponseTournament(tournament)

	return c.Status(200).JSON(responseUser)

}

// DeleteTournament function to delete a user
func DeleteTournament(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tournament models.Tournament

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindTournament(id, &tournament)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&tournament).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted Tournament")
}
