package controllers

import "os"

func GetRunningAddr() string {
	return os.Getenv("RUNNING_ADDR")
}
