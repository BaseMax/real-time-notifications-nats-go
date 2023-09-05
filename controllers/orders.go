package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/BaseMax/real-time-notifications-nats-go/notifications"
	"github.com/BaseMax/real-time-notifications-nats-go/rabbitmq"
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

	user := helpers.GetLoggedinInfo(c)
	admin, dbErr := models.GetAdmin()
	if dbErr != nil {
		return &dbErr.HttpErr
	}

	activities := models.Activity{
		UserID: admin.ID,
		Title:  fmt.Sprintf("We have new order from %s with order_id=%d", user.Username, order.ID),
		Action: models.ACTION_NEW_ORDER,
	}
	if err := notifications.Notify(activities); err != nil {
		return &err.HTTPError
	}

	if rabbitmq.RestartChannel() != nil {
		return echo.ErrInternalServerError
	}
	if err := rabbitmq.EnqueueTask(order, rabbitmq.QUEUE_NAME_ORDERS); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, order)
}

func FetchOrder(c echo.Context) error {
	return GetModelByPreload[models.Order](c, "id", "Products")
}

func FetchAllOrders(c echo.Context) error {
	return GetAllModelsByPreload[models.Order](c, "*", "Products")
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
	return ProcessFirstQueuedTask[models.Order](c, rabbitmq.QUEUE_NAME_ORDERS, models.TASK_BROWSE, "Products")
}

func DoneFirstQueuedOrder(c echo.Context) error {
	return ProcessFirstQueuedTask[models.Order](c, rabbitmq.QUEUE_NAME_ORDERS, models.TASK_DONE, "Products")
}

func CancelFirstQueuedOrder(c echo.Context) error {
	return ProcessFirstQueuedTask[models.Order](c, rabbitmq.QUEUE_NAME_ORDERS, models.TASK_CANCELED, "Products")
}
