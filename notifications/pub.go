package notifications

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
)

func Notify(activity models.Activity) *NotifErr {
	if err := models.Create(&activity); err != nil {
		return &NotifErr{Err: err.Err, HTTPError: err.HttpErr}
	}

	subject := CreateSubject(activity.RecieverID)
	data, _ := json.Marshal(activity)
	if err := conn.Publish(subject, data); err != nil {
		return &NotifErr{Err: err, HTTPError: *echo.ErrInternalServerError}
	}

	return nil
}
