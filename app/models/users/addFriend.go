package users

import "time"

type AddFriend struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	UserId   int  `gorm:"not null" json:"userId"`
	User     User `gorm:"foreignkey:UserId"`
	FriendId int  `gorm:"not null" json:"friendId"`
	Friend   User `gorm:"foreignkey:FriendId"`
	IsFriend bool `gorm:"bool" json:"is_friend"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
