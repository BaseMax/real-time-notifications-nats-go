package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if !helpers.IsUserAdmin(c) {
			return echo.ErrUnauthorized
		}
		return next(c)
	})
}
