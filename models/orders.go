package models

type Order struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	User       User   `json:"-"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Status     string `gorm:"default:IN-PROGRESS" json:"status"`
	ProductIDs []uint `gorm:"-" json:"product_ids"`
}
