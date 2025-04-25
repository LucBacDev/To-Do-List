package todotag

import (
	"source-base-go/infrastructure/repository/util"
)

type Service struct {
	todoTagRepo TodoTagRepository
	verifier util.Verifier
}

func NewService(todoTagRepo TodoTagRepository, verifier util.Verifier) *Service {
	return &Service{
		todoTagRepo: todoTagRepo,
		verifier: verifier,
	}
}

func (s Service) CreateTodoTag(todoId int, tagId int) error {
	err := s.todoTagRepo.Create(todoId, tagId)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteTodoTag(todoId int, tagId int) error  {
	err := s.todoTagRepo.Delete(todoId, tagId)
	if err != nil {
		return err
	}
	return nil
}
