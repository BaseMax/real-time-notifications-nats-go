package models

type Order struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Status   string    `gorm:"not null" json:"status"`
	Products []Product `json:"products"`
}
