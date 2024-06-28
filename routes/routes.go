package routes

import (
	"api/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/orderbook", handlers.GetOrderBook)
	e.POST("/orderbook", handlers.SaveOrderBook)
	e.GET("/orderhistory", handlers.GetOrderHistory)
	e.POST("/order", handlers.SaveOrder)
}
