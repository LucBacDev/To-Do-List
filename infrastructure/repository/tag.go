package repository

import (
	"source-base-go/entity"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (r TagRepository) GetAllTag() ([]*entity.Tag, error) {
	var listTag []*entity.Tag
	err := r.db.Find(&listTag).Error
	if err != nil {
		return nil, err
	}

	return listTag, nil
}

func (r TagRepository) Detail(id int) (*entity.Tag, error) {
	var tag *entity.Tag
	err := r.db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (r TagRepository) Create(data *entity.Tag) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TagRepository) Update(id int, data *entity.Tag) error {
	err := r.db.Model(&entity.Tag{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TagRepository) Delete(id int) error {
	err := r.db.Delete(&entity.Tag{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
