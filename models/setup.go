package models

import "gorm.io/gorm"

var db *gorm.DB

func Init(externalDb *gorm.DB) {
	db = externalDb
}
