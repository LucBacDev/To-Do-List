package tag

import (
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/util"
)

type Service struct {
	tagRepo TagRepository
	verifier util.Verifier
}

func NewService(tagRepo TagRepository, verifier util.Verifier) *Service {
	return &Service{
		tagRepo: tagRepo,
		verifier: verifier,
	}
}

func (s Service) GetAllTag() ([]*entity.Tag, error) {
	return s.tagRepo.GetAllTag()
}

func (s Service) CreateTag(tag *entity.Tag) error {
	err := s.tagRepo.Create(tag)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetTagDetail(id int) (*entity.Tag, error) {
	tag, err := s.tagRepo.Detail(id)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (s Service) UpdateTag(id int, tag *entity.Tag) error {
	err := s.tagRepo.Update(id, tag)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteTag(id int) error {
	err := s.tagRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
