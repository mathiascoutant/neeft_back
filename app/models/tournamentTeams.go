package models

import (
	"time"
)

type TournamentTeams struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	TournamentId uint       `gorm:"not null" json:"tournamentId"`
	Tournament   Tournament `gorm:"foreignkey:TournamentId"`
	TeamId       uint       `gorm:"not null" json:"teamId"`
	Team         Team       `gorm:"foreignkey:TeamId"`
	Status       uint       `gorm:"not null" json:"status"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
