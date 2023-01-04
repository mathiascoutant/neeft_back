package tournaments

/**
 * @Author ANYARONKE
 */

import (
	"neeft_back/app/models/games"
	"neeft_back/app/models/users"
	"time"
)

type Tournament struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `gorm:"varchar(255)" json:"name"`
	Price      uint       `gorm:"int" json:"price"`
	GameId     int        `gorm:"not null" json:"gameId"`
	Game       games.Game `gorm:"foreignkey:GameId"`
	OwnerId    int        `gorm:"not null" json:"ownerId"`
	Owner      users.User `gorm:"foreignkey:OwnerId"`
	TeamsLimit uint       `gorm:"uint" json:"teamsLimit"`
	IsFinished bool       `gorm:"bool" json:"isFinished"`
	Address    string     `gorm:"varchar(255)" json:"address"`
	Mode       string     `gorm:"varchar(255)" json:"mode"`
	StartDate  time.Time  `gorm:"datetime nullable" json:"startDate"`
	EndDate    time.Time  `gorm:"datetime nullable" json:"endDate"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
