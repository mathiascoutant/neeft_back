package teams

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"neeft_back/app/models/users"
	"time"
)

type Team struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	UserId          int        `gorm:"not null" json:"createBy"`
	User            users.User `gorm:"foreignkey:UserId"`
	Name            string     `gorm:"varchar(255)" json:"name"`
	UserCount       uint       `gorm:"uint" json:"userCount"`
	GameName        string     `gorm:"varchar(255)" json:"gameName"`
	TournamentCount uint       `gorm:"uint" json:"tournamentCount "`
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      time.Time
}
