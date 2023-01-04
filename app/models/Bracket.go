package models

import (
	"time"
)

type Bracket struct {
	ID           uint       `gorm:"primaryKey"   json:"id" `
	TeamId1      uint       `gorm:"not null" json:"teamId1"`
	Team1        Team       `gorm:"foreignkey:TeamId1"`
	TeamId2      uint       `gorm:"not null" json:"TeamId2"`
	Team2        Team       `gorm:"foreignkey:TeamId2"`
	TournamentId uint       `gorm:"not null" json:"tournamentId"`
	Tournament   Tournament `gorm:"foreignkey:TournamentId"`
	WinTeamId    uint       `gorm:"not null" json:"winTeamId"`
	WinTeam      Team       `gorm:"foreignkey:WinTeamId"`
	StatusGame   uint       `gorm:"not null" json:"statusGame"`
	Round        uint       `gorm:"not null" json:"round"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
