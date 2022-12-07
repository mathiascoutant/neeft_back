package database

/**
 * @Author ANYARONKE Dare Samuel
 */

import (
	"gorm.io/gorm"
	"neeft_back/app/models/teams"
	"neeft_back/app/models/tournaments"
	"neeft_back/app/models/users"
)

// RunMigration : Run the migration to initialize the database
func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&users.User{},
		&teams.Team{},
		&teams.UsersTeam{},
		&tournaments.Tournament{},
		&tournaments.TournamentTeams{},
		&users.AddFriend{},
	)
	if err != nil {
		return
	}
}
