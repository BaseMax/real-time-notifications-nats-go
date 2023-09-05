package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
)

func AddProduct(c echo.Context) error {
	product, err := CreateRecordFromModel[models.Product](c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)
}

func FetchProduct(c echo.Context) error {
	return GetModel[models.Product](c, "id")
}

func FetchAllProducts(c echo.Context) error {
	return GetAllModels[models.Product](c, "*")
}

func FetchProductsOfUser(c echo.Context) error {
	id, err := strconv.Atoi("order_id")
	if err != nil {
		return echo.ErrBadRequest
	}
	return GetAllModels[models.Product](c, "*", "user_id = ?", id)
}

func EditProduct(c echo.Context) error {
	return EditModelById[models.Product](c, "id")
}

func DeleteProduct(c echo.Context) error {
	return DeleteModelById[models.Product](c, "id")
}
