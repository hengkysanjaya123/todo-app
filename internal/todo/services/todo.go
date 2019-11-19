package services

import (
	"fmt"

	"github.com/hengkysanjaya/todo-app/internal/todo/models"
	"github.com/hengkysanjaya/todo-app/internal/todo/repositories"
)

// IService .
type IService interface {
	Fetch() []models.Todo
}

type service struct {
	TodoRepo repositories.ITodoRepository
}

// NewTodoService .
func NewTodoService() IService {
	return &service{
		TodoRepo: repositories.NewTodoRepository(),
	}
}

func (s *service) Fetch() []models.Todo {
	fmt.Println("Service")

	return s.TodoRepo.All()
}
