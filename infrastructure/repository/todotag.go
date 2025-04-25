package repository

import (
	"source-base-go/entity"
	"gorm.io/gorm"
)

type TodoTagRepository struct {
	db *gorm.DB
}

func NewTodoTagRepository(db *gorm.DB) *TodoTagRepository {
	return &TodoTagRepository{
		db: db,
	}
}


func (r TodoTagRepository) Create(todoId int, tagId int) error {
	todoTag := &entity.TodoTag{
		TodoId: todoId,
		TagId:  tagId,
	}

	err := r.db.Create(&todoTag).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TodoTagRepository) Delete(todoId int, tagId int) error {
	err := r.db.
		Where("todo_id = ? AND tag_id = ?", todoId, tagId).
		Delete(&entity.TodoTag{}).Error

	if err != nil {
		return err
	}
	return nil
}