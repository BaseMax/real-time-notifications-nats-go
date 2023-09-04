package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/real-time-notifications-nats-go/controllers"
	"github.com/BaseMax/real-time-notifications-nats-go/helpers"
	"github.com/BaseMax/real-time-notifications-nats-go/middlewares"
)

func todo(c echo.Context) error { return nil }

func InitRoutes() *echo.Echo {
	e := echo.New()
	g := e.Group("/", echojwt.WithConfig(echojwt.Config{SigningKey: helpers.GetJwtSecret()}))

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	g.POST("refresh", controllers.Refresh)
	g.GET("users/:id", controllers.FetchUser, middlewares.IsAdmin)
	g.GET("users", controllers.FetchAllUsers, middlewares.IsAdmin)
	g.DELETE("users/:id", controllers.DeleteUser, middlewares.IsAdmin)
	g.PUT("users/:id", controllers.EditUser, middlewares.IsAdmin)

	g.GET("notifications", controllers.Notification)
	g.GET("activities", controllers.FetchRecordedActivities)
	g.POST("activities/seen", controllers.SeenAllNotifications)

	g.POST("products", controllers.AddProduct, middlewares.IsAdmin)
	g.GET("products/:id", controllers.FetchProduct)
	g.GET("products", controllers.FetchAllProducts)
	g.GET("products/:order_id/orders", controllers.FetchProductsOfUser)
	g.PUT("products/:id", controllers.EditProduct, middlewares.IsAdmin)
	g.DELETE("products/:id", controllers.DeleteProduct, middlewares.IsAdmin)

	g.POST("orders", controllers.AddOrder)
	g.GET("orders/:id", controllers.FetchOrder)
	g.GET("orders", controllers.FetchAllOrders)
	g.PUT("orders/:id", controllers.EditOrder)
	g.DELETE("orders/:id", controllers.DeleteOrder, middlewares.IsAdmin)
	g.GET("orders/:id/status", controllers.FetchOrderStatus)
	g.POST("orders/:id/cancel", controllers.CancelOrder)
	g.GET("orders/first", controllers.GetFirstQueuedOrder, middlewares.IsAdmin)
	g.POST("orders/first/done", controllers.DoneFirstQueuedOrder, middlewares.IsAdmin)
	g.POST("orders/first/cancel", controllers.CancelFirstQueuedOrder, middlewares.IsAdmin)

	g.POST("refunds/:order_id", todo)
	g.GET("refunds/:id", todo)
	g.GET("refunds", todo)
	g.PUT("refunds/:id", todo)
	g.DELETE("refunds/:id", todo)
	g.GET("refunds/:id/status", todo)
	g.POST("refunds/:id/cancel", todo)
	g.GET("refunds/first", todo)
	g.POST("refunds/first/done", todo)
	g.POST("refunds/first/cancel", todo)

	return e
}
