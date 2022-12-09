package teams

import (
	"neeft_back/app/models/users"
)

type UsersTeam struct {
	ID     uint       `json:"id" gorm:"primaryKey"`
	UserId uint       `gorm:"not null" json:"userId"`
	User   users.User `gorm:"foreignkey:UserId"`
	TeamId uint       `gorm:"not null" json:"teamId"`
	Team   Team       `gorm:"foreignkey:TeamId"`
	Status uint	      `gorm:"not null" json:"status"`
}
