package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("godotenv:", err)
	}

	// DB setup

	// Models setup

	// Start application
	r := InitRoutes()
	r.Logger.Fatal(r.Start(":8000"))
}
