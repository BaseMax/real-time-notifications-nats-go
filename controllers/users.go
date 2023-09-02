package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/BaseMax/real-time-notifications-nats-go/notifications"
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

	admin, herr := models.GetAdmin()
	if herr != nil {
		return &herr.HttpErr
	}
	activity := models.Activity{
		UserID: admin.ID,
		Title:  user.Username + " registred to system",
		Action: models.ACTION_REGISTER,
	}
	if err := notifications.Notify(activity); err != nil {
		return &err.HTTPError
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
	user := helpers.GetLoggedinInfo(c)
	bearer := helpers.CreateJwtToken(user.ID, user.Username)
	return c.JSON(http.StatusOK, map[string]any{"bearer": bearer})
}

func FetchUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	user, herr := models.FindById[models.User](uint(id))
	user.Password = ""
	if herr != nil {
		return &herr.HttpErr
	}
	return c.JSON(http.StatusOK, user)
}

func FetchAllUsers(c echo.Context) error {
	users, err := models.GetAll[models.User]("id, username")
	if err != nil {
		return &err.HttpErr
	}
	return c.JSON(http.StatusOK, users)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := models.DeleteById(uint(id), &models.User{}); err != nil {
		return &err.HttpErr
	}
	return nil
}

func EditUser(c echo.Context) error {
	return nil
}
