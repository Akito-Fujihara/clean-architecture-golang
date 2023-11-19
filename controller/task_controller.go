package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct{}

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (t *TaskController) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
func (t *TaskController) Create(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	log.Print("Task: ", task)

	// created, err := usecase.Create(&task)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, nil)
	// }

	return c.JSON(http.StatusOK, nil)
}
