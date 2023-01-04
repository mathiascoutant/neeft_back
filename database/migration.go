package database

/**
 * @Author ANYARONKE Dare Samuel
 */

import (
	"gorm.io/gorm"
	"neeft_back/app/models/games"
	"neeft_back/app/models/roles"
	"neeft_back/app/models/teams"
	"neeft_back/app/models/tournaments"
	"neeft_back/app/models/users"
)

// RunMigration : Run the migration to initialize the database
func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		// Users
		&roles.Role{},
		&users.User{},
		&users.UserROle{},
		&teams.Team{},
		&teams.InfoPro{},
		&teams.UsersTeam{},
		&games.Game{},
		&tournaments.Tournament{},
		&tournaments.TeamRegistrationTournament{},
		&tournaments.TournamentPlayer{},
		&tournaments.Bracket{},
		&users.AddFriend{},
	)
	if err != nil {
		return
	}
}
