package models

type Product struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	OrderID     uint
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"not null" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
}
