package models

import (
	"time"
)

type UserRole struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserId int  `gorm:"not null" json:"userId"`
	User   User `gorm:"foreignkey:UserId"`
	RoleId int  `gorm:"not null" json:"roleId"`
	Role   Role `gorm:"foreignkey:RoleId"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
