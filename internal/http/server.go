package http

import (
	"github.com/go-chi/chi"
	"github.com/hengkysanjaya/todo-app/internal/http/controller"
)

// CompileRouter bla bla
func CompileRouter() chi.Router {
	todoController := controller.NewTodoController()
	_ = controller.NewReminderController()

	r := chi.NewRouter()

	r.Get("/todos", todoController.Get)

	return r
}
