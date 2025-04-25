package tag

import (
	"source-base-go/entity"
)

type TagRepository interface {
	GetAllTag() ([]*entity.Tag, error)
	Detail(id int) (*entity.Tag, error)
	Create(data *entity.Tag) error
	Update(id int, data *entity.Tag) error
	Delete(id int) error
}

type UseCase interface {
	GetAllTag() ([]*entity.Tag, error)
	GetTagDetail(id int) (*entity.Tag, error)
	CreateTag(tag *entity.Tag) error
	UpdateTag(id int, tag *entity.Tag) error
	DeleteTag(id int) error
}
