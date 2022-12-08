package tournaments

import (
	"neeft_back/app/models/teams"
)

const (
	StatusPending  = 0
	StatusAccepted = 1
	StatusRejected = 2
)

type TournamentTeams struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	TournamentId uint       `gorm:"not null" json:"tournamentId"`
	Tournament   Tournament `gorm:"foreignkey:TournamentId"`
	TeamId       uint       `gorm:"not null" json:"teamId"`
	Team         teams.Team `gorm:"foreignkey:TeamId"`
	Status       uint       `gorm:"not null" json:"status"`
}
