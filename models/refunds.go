package models

type Refund struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	OrderID uint   `gorm:"not null; unique" json:"order_id"`
	Status  string `gorm:"not null;" json:"status"`
	Order   Order  `json:"orders"`
}
