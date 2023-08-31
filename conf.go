package main

import "os"

func GetRunningAddr() string {
	return os.Getenv("RUNNING_ADDR")
}

func GetJwtSecret() string {
	return os.Getenv("RUNNING_SECRET")
}
