package models

import (
	"time"
)

type UsersTeam struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserId uint `gorm:"not null" json:"userId"`
	User   User `gorm:"foreignkey:UserId"`
	TeamId uint `gorm:"not null" json:"teamId"`
	Team   Team `gorm:"foreignkey:TeamId"`
	Status uint `gorm:"not null" json:"status"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
