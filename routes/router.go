package routes

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/controllers/authController"
	"neeft_back/app/controllers/teams"
	"neeft_back/app/controllers/tournament"
	"neeft_back/app/controllers/users"
	"neeft_back/middleware"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")
	api.Post("login", authController.Login)
	api.Post("/user", users.CreateUser)
	api.Post("/register", users.CreateUser)

	auth := app.Use(middleware.VerifyJWT)

	//------------------ Users ------------------
	auth.Get("/users", users.GetAllUser)
	auth.Get("/user/:id", users.GetUser)
	auth.Put("/user/:id", users.UpdateUser)
	auth.Delete("/user/:id", users.DeleteUser)

	//------------------ Users Friend ------------------
	auth.Post("/friend", users.CreateUserFriend)
	auth.Get("/show-friend/:id", users.GetUserFriends)

	//------------------ Teams ------------------
	auth.Post("/team", teams.CreateTeam)
	auth.Get("/teams", teams.GetAllTeam)
	auth.Get("/team/:id", teams.GetTeam)

	//------------------ Tournaments ------------------
	auth.Post("/tournament", tournament.CreateTournament)
	auth.Get("/tournaments", tournament.GetAllTournament)
	auth.Get("/tournament/:id", tournament.GetTournament)
	auth.Delete("/tournament/:id", tournament.DeleteTournament)

	// Debug
	// api.Get("/jwt/debug", users.JWTDebug)
}
