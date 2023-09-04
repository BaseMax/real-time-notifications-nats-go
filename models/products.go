package models

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	OrderID     uint    `json:"order_id"`
	Title       string  `gorm:"not null" json:"title"`
	Description string  `gorm:"not null" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
}

func ReserveProducts(orderId uint, productIds []uint) *DbErr {
	r := db.Model(Product{}).
		Where("order_id = 0 AND id IN ?", productIds).
		Update("order_id", orderId)
	return errGormToHttp(r)
}
