package badges

const (
	RecipientTeam uint = 0
	RecipientUser uint = 1
)

type Badge struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Recipient   uint   `json:"recipient" gorm:"not null"`
	RecipientId uint   `json:"recipient_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Section     uint   `json:"section" gorm:"not null"`
	Category    uint   `json:"category" gorm:"not null"`
}

type BadgeInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Section     uint   `json:"section"`
	Category    uint   `json:"category"`
}
