package models

import (
	"time"
)

type TeamRegistrationTournament struct {
	ID                 uint       `gorm:"primaryKey"   json:"id" `
	TeamId             uint       `gorm:"not null" json:"teamId"`
	Team               Team       `gorm:"foreignkey:TeamId"`
	TournamentId       uint       `gorm:"not null" json:"tournamentId"`
	Tournament         Tournament `gorm:"foreignkey:TournamentId"`
	RegistrationStatue uint       `gorm:"not null" json:"registrationStatue"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
