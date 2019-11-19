package controller

import (
	"fmt"
	"github.com/hengkysanjaya/todo-app/internal/todo/services"
	"net/http"
)

// IReminderController .
type IReminderController interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type reminderController struct {
	TodoService services.IService
}

// NewReminderController .
func NewReminderController() IReminderController {
	return &reminderController{
		TodoService: services.NewTodoService(),
	}
}

// Get .
func (c *reminderController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Controller")

	c.TodoService.Fetch()
}
