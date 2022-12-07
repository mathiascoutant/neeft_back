package tournaments

/**
 * @Author ANYARONKE
 */

import (
	"time"
)

type Tournament struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `gorm:"varchar(255)" json:"name"`
	Count      uint   `gorm:"uint" json:"count"`
	Price      uint   `gorm:"int" json:"price"`
	Game       string `gorm:"varchar(255)" json:"game"`
	TeamsCount uint   `gorm:"uint" json:"teamsCount"`
	IsFinished bool   `gorm:"bool" json:"isFinished"`
	Mode       string `gorm:"varchar(255)" json:"mode"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
