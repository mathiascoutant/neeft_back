package users

import (
	"neeft_back/app/models/roles"
	"time"
)

type UserROle struct {
	ID     uint       `json:"id" gorm:"primaryKey"`
	UserId int        `gorm:"not null" json:"userId"`
	User   User       `gorm:"foreignkey:UserId"`
	RoleId int        `gorm:"not null" json:"roleId"`
	Role   roles.Role `gorm:"foreignkey:RoleId"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
