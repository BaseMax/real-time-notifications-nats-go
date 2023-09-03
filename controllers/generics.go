package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/labstack/echo/v4"
)

func GetModel[T any](c echo.Context, idParam string) error {
	id, err := strconv.Atoi(c.Param(idParam))
	if err != nil {
		return echo.ErrBadRequest
	}
	model, dbErr := models.FindById[T](uint(id))
	if dbErr != nil {
		return &dbErr.HttpErr
	}
	fmt.Println(model)
	return c.JSON(http.StatusOK, model)
}
