package database

import (
	"fmt"
	"os"
	"strconv"
)

type DbConf struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint
	TimeZone string
}

func ReadConfig() (*DbConf, error) {
	host := os.Getenv("POSTGRES_HOSTNAME")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	timeZone := os.Getenv("POSTGRES_TIMEZONE")

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return nil, fmt.Errorf("postgres port: %w", err)
	}

	return &DbConf{host, user, password, dbName, uint(port), timeZone}, nil
}
