package tournament

/**
 * @Author ANYARONKE
 */

import (
	"github.com/gofiber/fiber/v2"
	teams2 "neeft_back/app/controllers/teams"
	"neeft_back/app/models/teams"
	"neeft_back/app/models/tournaments"
	"neeft_back/database"
)

type TournamentTeamsRequest struct {
	TeamId uint `json:"teamId"`
	Status uint `json:"status"`
}

// SendPendingRequest is used to send a pending request to join a tournament by a team
func SendPendingRequest(c *fiber.Ctx) error {
	var request TournamentTeamsRequest
	id, _ := c.ParamsInt("id")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var tournament tournaments.Tournament
	err := FindTournament(id, &tournament)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var team teams.Team
	err = teams2.FindTeam(id, &team)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	finalItem := tournaments.TournamentTeams{
		TournamentId: uint(id),
		TeamId:       request.TeamId,
		Status:       0,
		Tournament:   tournament,
		Team:         team,
	}

	database.Database.Db.Create(&finalItem)
	c.Status(200)
	return nil
}
