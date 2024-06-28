package main

import (
	"api/config"
	"api/db"
	"api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db.Init()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
