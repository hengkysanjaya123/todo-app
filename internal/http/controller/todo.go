package controller

import (
	"fmt"
	"net/http"

	libHttp "github.com/hengkysanjaya/todo-app/lib/http"

	"github.com/hengkysanjaya/todo-app/internal/todo/services"
)

// ITodoController .
type ITodoController interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	TodoService services.IService
}

// NewTodoController .
func NewTodoController() ITodoController {
	return &todoController{
		TodoService: services.NewTodoService(),
	}
}

// Get .
func (c *todoController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Controller")

	todos := c.TodoService.Fetch()

	libHttp.ResponseJSON(w, todos, 200)
}
