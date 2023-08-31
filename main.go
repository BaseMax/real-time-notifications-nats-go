package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/BaseMax/real-time-notifications-nats-go/controllers"
	"github.com/BaseMax/real-time-notifications-nats-go/database"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("godotenv: ", err)
	}

	c, err := database.ReadConfig()
	if err != nil {
		log.Fatal("readconfig: ", err)
	}
	conn, err := database.OpenPostgres(c)
	if err != nil {
		log.Fatal("open postgres: ", err)
	}

	if err := models.Init(conn); err != nil {
		log.Fatal("models init: ", err)
	}

	r := InitRoutes()
	r.Logger.Fatal(r.Start(controllers.GetRunningAddr()))
}
