package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenPostgres(c *DbConf) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		c.Host, c.User, c.Password, c.DbName, c.Port, c.TimeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard, TranslateError: true,
	})
	return
}
