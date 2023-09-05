package models

type Refund struct {
	ID      uint   `gorm:"primaryKey" json:"id,omitempty"`
	OrderID uint   `gorm:"not null; unique" json:"order_id,omitempty"`
	Status  string `gorm:"default:IN-PROGRESS" json:"status,omitempty"`
	Order   Order  `json:"order,omitempty"`
}

func (r Refund) GetID() uint {
	return r.ID
}

func (r Refund) GetStatus() string {
	return r.Status
}
