package routes

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/controllers/authController"
	"neeft_back/app/controllers/teams"
	"neeft_back/app/controllers/tournament"
	"neeft_back/app/controllers/users"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")
	api.Post("login", authController.Login)

	//------------------ Users ------------------
	api.Post("/user", users.CreateUser)
	api.Get("/users", users.GetAllUser)
	api.Get("/user/:id", users.GetUser)
	api.Put("/user/:id", users.UpdateUser)
	api.Delete("/user/:id", users.DeleteUser)

	//------------------ Users Friend ------------------
	api.Post("/friend", users.CreateUserFriend)
	api.Get("/show-friend/:id", users.GetUserFriends)

	//------------------ Teams ------------------
	api.Post("/team", teams.CreateTeam)
	api.Get("/teams", teams.GetAllTeam)
	api.Get("/team/:id", teams.GetTeam)

	//------------------ Tournaments ------------------
	api.Post("/tournament", tournament.CreateTournament)
	api.Get("/tournaments", tournament.GetAllTournament)
	api.Get("/tournament/:id", tournament.GetTournament)
	api.Delete("/tournament/:id", tournament.DeleteTournament)

}
