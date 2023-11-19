package main

import (
	"github.com/Akito-Fujihara/clean-architecture-golang/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	TaskController := controller.TaskController{}

	e.GET("/tasks", TaskController.Get)
	e.POST("/tasks", TaskController.Create)

	e.Logger.Fatal(e.Start(":8080"))
}
