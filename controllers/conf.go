package controllers

import "os"

func GetRunningAddr() string {
	return os.Getenv("RUNNING_ADDR")
}

func GetJwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
