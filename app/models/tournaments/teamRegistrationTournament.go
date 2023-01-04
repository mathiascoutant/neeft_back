package tournaments

import (
	"neeft_back/app/models/teams"
	"time"
)

type TeamRegistrationTournament struct {
	ID                 uint       `gorm:"primaryKey"   json:"id" `
	TeamId             uint       `gorm:"not null" json:"teamId"`
	Team               teams.Team `gorm:"foreignkey:TeamId"`
	TournamentId       uint       `gorm:"not null" json:"tournamentId"`
	Tournament         Tournament `gorm:"foreignkey:TournamentId"`
	RegistrationStatue uint       `gorm:"not null" json:"registrationStatue"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
