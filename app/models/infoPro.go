package models

import (
	"time"
)

type InfoPro struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `gorm:"varchar(255)" json:"name"`
	TeamId uint   `gorm:"not null" json:"teamId"`
	Team   Team   `gorm:"foreignkey:TeamId"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
