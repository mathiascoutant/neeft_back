package models

/**
 * @Author ANYARONKE Daré Samuel
 */

import "time"

type User struct {
	ID              uint   `gorm:"primaryKey"   json:"id" `
	Username        string `gorm:"varchar(255)" json:"username"`
	FirstName       string `gorm:"varchar(255)" json:"firstName"`
	LastName        string `gorm:"varchar(255)" json:"lastName"`
	Email           string `gorm:"varchar(255)" json:"email"`
	EmailVerifiedAt bool   `gorm:"boolean"      json:"emailVerifiedAt"`
	Password        string `gorm:"varchar(255)" json:"password"`
	RememberToken   string `gorm:"varchar(100)" json:"rememberToken"`
	BirthDate       string `gorm:"varchar(255)" json:"birthDate"`
	Avatar          string `gorm:"varchar(255)" json:"avatar"`
	IsBan           bool   `gorm:"boolean"      json:"isBan"`
	LastLoginAt     time.Time
	IsSuperAdmin    bool `gorm:"boolean"      json:"isSuperAdmin"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
