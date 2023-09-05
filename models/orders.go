package models

type Order struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	User     User      `json:"-"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	Status   string    `gorm:"default:IN-PROGRESS" json:"status"`
	Products []Product `gorm:"many2many:order_products" json:"products,omitempty"`
}

func (o Order) GetID() uint {
	return o.ID
}

func (o Order) GetStatus() string {
	return o.Status
}

func (o Order) GetOwnerID() uint {
	return o.UserID
}

func FetchOrder(id uint) (order *Order, err *DbErr) {
	r := db.Preload("Products").First(&order, id)
	return order, errGormToHttp(r)
}

func FetchAllOrders() (orders *[]Order, err *DbErr) {
	r := db.Preload("Products").Find(&orders)
	return orders, errGormToHttp(r)
}
