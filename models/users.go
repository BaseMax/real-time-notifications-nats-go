package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null; unique" json:"username"`
	Password string `gorm:"not null" json:"password,omitempty"`
}

func Login(u *User) *DbErr {
	r := db.Where(&u).First(&u)
	return errGetGormToHttp(r)
}

func GetAdmin() (user *User, err *DbErr) {
	r := db.First(&user, "username = ?", GetAdminConf().Username)
	err = errGetGormToHttp(r)
	return
}
