package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BaseMax/real-time-notifications-nats-go/models"
	"github.com/labstack/echo/v4"
)

func AddProduct(c echo.Context) error {
	var product models.Product
	if err := json.NewDecoder(c.Request().Body).Decode(&product); err != nil {
		return echo.ErrBadRequest
	}
	if err := models.Create(&product); err != nil {
		return &err.HttpErr
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
	return nil
}

func DeleteProduct(c echo.Context) error {
	return DeleteModelById[models.Product](c, "id")
}
