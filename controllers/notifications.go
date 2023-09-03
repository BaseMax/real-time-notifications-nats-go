package controllers

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
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

	conn := notifications.GetConn()
	subject := notifications.CreateSubject(helpers.GetLoggedinInfo(c).ID)
	sub, err = conn.Subscribe(subject, func(msg *nats.Msg) {
		json.Unmarshal(msg.Data, &data)
		err := ws.WriteJSON(data)
		if err != nil {
			cleaner()
		}
	})
	if err != nil {
		return echo.ErrInternalServerError
	}

	for {
		select {}
	}
}
