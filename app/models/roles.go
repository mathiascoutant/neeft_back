package models

import "time"

type Role struct {
	ID          uint   `gorm:"primaryKey"   json:"id" `
	Name        string `gorm:"varchar(255)" json:"name"`
	Description string `gorm:"varchar(255)" json:"description"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
