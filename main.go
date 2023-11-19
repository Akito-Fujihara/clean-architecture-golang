package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", func(c echo.Context) error {
		return c.String(200, "Get all tasks")
	})
	e.POST("/tasks", func(c echo.Context) error {
		return c.String(200, "Create task")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
