package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/BaseMax/real-time-notifications-nats-go/rabbitmq"
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
	return c.JSON(http.StatusOK, model)
}

func GetModelByPreload[T any](c echo.Context, idParam, preload string) error {
	id, err := strconv.Atoi(c.Param(idParam))
	if err != nil {
		return echo.ErrBadRequest
	}
	model, dbErr := models.FindByIdPreload[T](uint(id), preload)
	if dbErr != nil {
		return &dbErr.HttpErr
	}
	return c.JSON(http.StatusOK, model)
}

func GetAllModels[T any](c echo.Context, sel string, con ...any) error {
	models, dbErr := models.GetAll[T](sel, con...)
	if dbErr != nil {
		return &dbErr.HttpErr
	}
	return c.JSON(http.StatusOK, models)
}

func GetAllModelsByPreload[T any](c echo.Context, sel, preload string, cond ...any) error {
	allModels, err := models.GetAllPreload[T](sel, preload, cond...)
	if err != nil {
		return &err.HttpErr
	}
	return c.JSON(http.StatusOK, allModels)
}

func EditModelById[T any](c echo.Context, idParam string) error {
	var model T
	id, err := strconv.Atoi(c.Param(idParam))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&model); err != nil {
		return echo.ErrBadRequest
	}
	if err := models.UpdateById(uint(id), &model); err != nil {
		return &err.HttpErr
	}
	return c.NoContent(http.StatusOK)
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

func CreateRecordFromModel[T any](c echo.Context) (*T, error) {
	var model T
	if err := json.NewDecoder(c.Request().Body).Decode(&model); err != nil {
		return nil, echo.ErrBadRequest
	}
	if err := models.Create(&model); err != nil {
		return nil, &err.HttpErr
	}
	return &model, nil
}

func ProcessFirstQueuedTask[T any](c echo.Context, queueName string, newStatus, preload string) error {
	if rabbitmq.RestartChannel() != nil {
		return echo.ErrInternalServerError
	}

	model, err := rabbitmq.ProcessFirstTask[T](queueName, newStatus, preload)
	if err != nil {
		return echo.ErrInternalServerError
	}
	if model == nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, model)
}
