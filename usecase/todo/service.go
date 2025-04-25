package todo

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
)

type Service struct {
	todoRepo TodoRepository
	verifier util.Verifier
}

func NewService(todoRepo TodoRepository, verifier util.Verifier) *Service {
	return &Service{
		todoRepo: todoRepo,
		verifier: verifier,
	}
}

func (s Service) GetAllTodo() ([]*entity.Todo, error) {
	return s.todoRepo.GetAllTodo()
}

func (s Service) CreateTodo(todo *entity.Todo) error {
	err := s.todoRepo.Create(todo)
	if err != nil {
		return err
	}
	return nil
}
func (s Service) GetTodoDetail(id int) (*entity.Todo, error) {
	todo, err := s.todoRepo.Detail(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s Service) UpdateTodo(id int, todo *entity.Todo) error {
	err := s.todoRepo.Update(id, todo)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteTodo(id int) error {
	err := s.todoRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
