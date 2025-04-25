package entity

type TodoTag struct {
	TodoId int `gorm:"column:todo_id;" json:"todo_id"`
	TagId  int `gorm:"column:tag_id;" json:"tag_id"`
}
