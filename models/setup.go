package models

import "gorm.io/gorm"

var db *gorm.DB

func Init(externalDb *gorm.DB) error {
	db = externalDb
	return db.AutoMigrate(&User{}, &Activity{}, &Product{}, &Order{}, &Refund{})
}
