package entity

import (
	"time"
)

type Priority string

const (
	LowPriority    Priority = "low"
	MediumPriority Priority = "medium"
	HighPriority   Priority = "high"
)

type Todo struct {
	Id          int    `gorm:"column:id;primaryKey" json:"id"`
	UserId      int    `gorm:"column:user_id" json:"user_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	IsDone      bool      `gorm:"column:is_done" json:"is_done"`
	Priority    Priority  `gorm:"column:priority" json:"priority"`
	DueDate     *time.Time `gorm:"column:due_date" json:"due_date"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
