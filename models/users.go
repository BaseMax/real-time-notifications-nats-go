package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null; unique" json:"username"`
	Password string `gorm:"not null" json:"password,omitempty"`
}

func (o User) GetID() uint {
	return o.ID
}

func (o User) GetStatus() string {
	return ""
}

func (o User) GetOwnerID() uint {
	return o.ID
}

func (User) GetName() string {
	return "user"
}

func Login(u *User) *DbErr {
	r := db.Where(&u).First(&u)
	return errGormToHttp(r)
}

func GetAdmin() (user *User, err *DbErr) {
	r := db.First(&user, "username = ?", GetAdminConf().Username)
	err = errGormToHttp(r)
	return
}
