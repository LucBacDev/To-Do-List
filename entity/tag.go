package entity
import (
	"time"
)

type Tag struct {
	Id    int  `gorm:"column:id;primaryKey" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Color string `gorm:"column:color" json:"color"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`

}
