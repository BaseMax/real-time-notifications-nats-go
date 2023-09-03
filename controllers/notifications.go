package controllers

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/BaseMax/real-time-notifications-nats-go/notifications"
)

var upgrader = websocket.Upgrader{}

func Notification(c echo.Context) error {
	var sub *nats.Subscription
	var data map[string]any

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	cleaner := func() {
		sub.Unsubscribe()
		ws.Close()
	}
	defer cleaner()

	userId := helpers.GetLoggedinInfo(c).ID
	notifs, dbErr := models.GetActivitiesByUserId(userId)
	if dbErr != nil && dbErr.Err != nil {
		return &dbErr.HttpErr
	}

	if len(*notifs) > 0 {
		if err := ws.WriteJSON(notifs); err != nil {
			return echo.ErrBadRequest
		}
		if dbErr := models.SeenActivities(userId); dbErr != nil {
			return &dbErr.HttpErr
		}
	}

	conn := notifications.GetConn()
	subject := notifications.CreateSubject(userId)
	sub, err = conn.Subscribe(subject, func(msg *nats.Msg) {
		json.Unmarshal(msg.Data, &data)
		err := ws.WriteJSON(data)
		if err != nil {
			cleaner()
		}

		models.SeenActivities(userId)
	})
	if err != nil {
		return echo.ErrInternalServerError
	}

	for {
		select {}
	}
}
