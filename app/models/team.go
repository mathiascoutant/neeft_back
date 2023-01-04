package models

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"time"
)

type Team struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `gorm:"varchar(255)" json:"name"`
	Description string `gorm:"varchar(255)" json:"description"`
	OwnerId     int    `gorm:"not null" json:"ownerId"`
	Owner       User   `gorm:"foreignkey:OwnerId"`
	Type        string `gorm:"varchar(255)" json:"type"`
	IsBanned    bool   `gorm:"not null default:false" json:"isBanned"`
	MaxMembers  uint   `gorm:"uint" json:"maxMembers"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
