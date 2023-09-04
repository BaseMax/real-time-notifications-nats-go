package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/labstack/echo/v4"
)

func AddOrder(c echo.Context) error {
	order, err := CreateRecordFromModel[models.Order](c)
	// GORM postgres driver doesn't have gorm.ErrForeignKeyViolated translation
	// I should hack
	if errors.Is(err, echo.ErrInternalServerError) {
		return echo.ErrNotFound
	}
	if err != nil {
		return err
	}

	if err := models.ReserveProducts(order.ID, order.ProductIDs); err != nil {
		return &err.HttpErr
	}

	// Notify Admin

	// Queue order

	return c.JSON(http.StatusOK, order)
}

func FetchOrder(c echo.Context) error {
	return GetModel[models.Order](c, "id")
}

func FetchAllOrders(c echo.Context) error {
	return GetAllModels[models.Order](c, "*")
}

func EditOrder(c echo.Context) error {
	return EditModelById[models.Order](c, "id")
}

func DeleteOrder(c echo.Context) error {
	return DeleteModelById[models.Order](c, "id")
}

func FetchOrderStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	order, dbErr := models.FindById[models.Order](uint(id))
	if dbErr != nil {
		return &dbErr.HttpErr
	}
	return c.JSON(http.StatusOK, map[string]any{"status": order.Status})
}

func CancelOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	dbErr := models.UpdateById[models.Order](uint(id), &models.Order{Status: models.TASK_CANCELED})
	if dbErr != nil {
		return &dbErr.HttpErr
	}
	return c.JSON(http.StatusOK, map[string]any{"status": models.TASK_CANCELED})
}

func GetFirstQueuedOrder(c echo.Context) error {
	return nil
}

func DoneFirstQueuedOrder(c echo.Context) error {
	return nil
}

func CancelFirstQueuedOrder(c echo.Context) error {
	return nil
}
