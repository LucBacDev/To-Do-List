package todo

import (
	"source-base-go/entity"
)

type TodoRepository interface {
	GetAllTodo() ([]*entity.Todo, error)
	Detail(id int) (*entity.Todo, error)
	Create(data *entity.Todo) error
	Update(id int, data *entity.Todo) error
	Delete(id int) error
}

type UseCase interface {
	GetAllTodo() ([]*entity.Todo, error)
	GetTodoDetail(id int) (*entity.Todo, error)
	CreateTodo(Todo *entity.Todo) error
	UpdateTodo(id int, Todo *entity.Todo) error
	DeleteTodo(id int) error
	
}
