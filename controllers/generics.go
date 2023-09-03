package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
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

func GetAllModels[T any](c echo.Context, sel string, con ...any) error {
	models, dbErr := models.GetAll[T](sel, con...)
	if dbErr != nil {
		return &dbErr.HttpErr
	}
	return c.JSON(http.StatusOK, models)
}

func DeleteModelById[T any](c echo.Context, idParam string) error {
	id, err := strconv.Atoi(c.Param(idParam))
	if err != nil {
		return echo.ErrBadRequest
	}
	var model T
	if err := models.DeleteById(uint(id), &model); err != nil {
		return &err.HttpErr
	}
	return c.JSON(http.StatusOK, model)
}
