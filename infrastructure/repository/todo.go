package repository

import (
	"source-base-go/entity"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r TodoRepository) GetAllTodo() ([]*entity.Todo, error) {
	var listTodo []*entity.Todo
	err := r.db.Find(&listTodo).Error
	if err != nil {
		return nil, err
	}

	return listTodo, nil
}

func (r TodoRepository) Detail(id int) (*entity.Todo, error) {
	var listTodo *entity.Todo
	err := r.db.Where("id = ?", id).Find(&listTodo).Error
	if err != nil {
		return nil, err
	}

	return listTodo, nil
}

func (r TodoRepository) Create(data *entity.Todo) error {
	err := r.db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TodoRepository) Update(id int, data *entity.Todo) error {
	err := r.db.Model(&entity.Todo{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TodoRepository) Delete(id int) error {
	err := r.db.Delete(&entity.Todo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
