package controllers

import (
	"errors"
	"net/http"

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

	// Queue order

	return c.JSON(http.StatusOK, order)
}

func FetchOrder(c echo.Context) error {
	return nil
}

func FetchAllOrders(c echo.Context) error {
	return nil
}

func EditOrder(c echo.Context) error {
	return nil
}

func DeleteOrder(c echo.Context) error {
	return nil
}

func FetchOrderStatus(c echo.Context) error {
	return nil
}

func CancelOrder(c echo.Context) error {
	return nil
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
