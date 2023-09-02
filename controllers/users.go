package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
)

func Register(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	user.Password = models.HashPassword(user.Password)
	if err := models.Create(&user); err != nil {
		return &err.HttpErr
	}
	bearer := helpers.CreateJwtToken(user.ID, user.Username)
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}

func Login(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	user.Password = models.HashPassword(user.Password)
	if err := models.Login(&user); err != nil {
		return &err.HttpErr
	}
	bearer := helpers.CreateJwtToken(user.ID, user.Username)
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}

func Refresh(c echo.Context) error {
	bearer := helpers.CreateJwtToken(helpers.GetLoggedinInfo(c))
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}

func FetchUser(c echo.Context) error {
	
	return nil
}

func FetchAllUsers(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	return nil
}

func EditUser(c echo.Context) error {
	return nil
}
