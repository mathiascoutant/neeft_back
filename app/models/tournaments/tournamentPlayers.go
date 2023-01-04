package tournaments

import (
	"neeft_back/app/models/teams"
	"neeft_back/app/models/users"
	"time"
)

type TournamentPlayer struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	TournamentId uint       `gorm:"not null" json:"tournamentId"`
	Tournament   Tournament `gorm:"foreignkey:TournamentId"`
	TeamId       uint       `gorm:"not null" json:"teamId"`
	Team         teams.Team `gorm:"foreignkey:TeamId"`
	UserId       uint       `gorm:"not null" json:"userId"`
	User         users.User `gorm:"foreignkey:UserId"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
