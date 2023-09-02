package models

import (
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null; unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

func Login(u *User) *DbErr {
	err := db.Where(&u).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &DbErr{Err: err, HttpErr: *echo.ErrNotFound}
	} else if err != nil {
		return &DbErr{Err: err, HttpErr: *echo.ErrInternalServerError}
	}
	return nil
}
