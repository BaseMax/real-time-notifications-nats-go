package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
)

func FetchRecordedActivities(c echo.Context) error {
	activities, err := models.GetActivitiesByUserId(helpers.GetLoggedinInfo(c).ID)
	if err != nil {
		return &err.HttpErr
	}
	return c.JSON(http.StatusOK, activities)
}

func SeenAllNotifications(c echo.Context) error {
	err := models.SeenActivities(helpers.GetLoggedinInfo(c).ID)
	if err != nil {
		return &err.HttpErr
	}
	return c.NoContent(http.StatusNoContent)
}
