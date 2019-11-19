package repositories

import "fmt"

import "github.com/hengkysanjaya/todo-app/internal/todo/models"

// ITodoRepository .
type ITodoRepository interface {
	All() []models.Todo
	Insert()
	Update()
	Delete()
}

type todoRepository struct {
}

// NewTodoRepository .
func NewTodoRepository() ITodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) All() []models.Todo {
	fmt.Println("REPO - ALL")

	return []models.Todo{
		models.Todo{
			Id:          1,
			Title:       "A",
			Description: "AAA",
		},
		{
			Id:          2,
			Title:       "B",
			Description: "BBB",
		},
	}
}

func (tr *todoRepository) Insert() {

}

func (tr *todoRepository) Update() {

}

func (tr *todoRepository) Delete() {

}
