package models

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id,omitempty"`
	OrderID     uint    `json:"order_id,omitempty"`
	Title       string  `gorm:"not null" json:"title,omitempty"`
	Description string  `gorm:"not null" json:"description,omitempty"`
	Price       float64 `gorm:"not null" json:"price,omitempty"`
	Orders      []Order `gorm:"many2many:order_products" json:"orders,omitempty"`
}
