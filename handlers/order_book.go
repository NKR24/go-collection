package handlers

import (
	"api/db"
	"api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetOrderBook(c echo.Context) error {
	exchangeName := c.QueryParam("exchange_name")
	pair := c.QueryParam("pair")

	var orderBook models.OrderBook
	if err := db.DB.Where("exchange = ? AND pair = ?", exchangeName, pair).First(&orderBook).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order book not found"})
	}

	return c.JSON(http.StatusOK, orderBook)
}

func SaveOrderBook(c echo.Context) error {
	exchangeName := c.QueryParam("exchange_name")
	pair := c.QueryParam("pair")

	var orderBook models.OrderBook
	if err := c.Bind(&orderBook); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	orderBook.Exchange = exchangeName
	orderBook.Pair = pair

	if err := db.DB.Save(&orderBook).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save order book"})
	}

	return c.JSON(http.StatusOK, orderBook)
}
