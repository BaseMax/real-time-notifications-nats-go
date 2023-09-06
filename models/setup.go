package models

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

	"gorm.io/gorm"
)

var db *gorm.DB

func HashPassword(pass string) string {
	hashByte := sha256.Sum256([]byte(pass))
	hashStr := hex.EncodeToString(hashByte[:])
	return hashStr
}

func GetAdminConf() User {
	return User{Username: os.Getenv("ADMIN_USERNAME"), Password: os.Getenv("ADMIN_PASSWORD")}
}

func Init(externalDb *gorm.DB) error {
	db = externalDb
	if err := db.AutoMigrate(&User{}, &Activity{}, &Order{}, &Product{}, &Refund{}); err != nil {
		return err
	}

	admin := GetAdminConf()
	admin.Password = HashPassword(admin.Password)

	var count int64
	db.Where(&admin).First(&User{}).Count(&count)
	if count == 1 {
		return nil
	}
	return db.Create(&admin).Error
}
