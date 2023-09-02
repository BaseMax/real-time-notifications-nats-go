package middlewares

import (
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/models"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		_, loggeninName := helpers.GetLoggedinInfo(c)
		user := models.GetAdminConf()
		if loggeninName != user.Username {
			return echo.ErrUnauthorized
		}
		return next(c)
	})
}
