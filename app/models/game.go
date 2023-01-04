package models

type Game struct {
	ID          uint   `gorm:"primaryKey"   json:"id" `
	Name        string `gorm:"varchar(255)" json:"name"`
	Description string `gorm:"varchar(255)" json:"description"`
	PlayerLimit uint   `gorm:"uint" json:"playerLimit"`
}
