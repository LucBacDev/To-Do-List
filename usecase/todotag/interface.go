package todotag


type TodoTagRepository interface {
	Create(todoId int, tagId int) error
	Delete(todoId int, tagId int) error
}

type UseCase interface {
	CreateTodoTag(todoId int, tagId int) error
	DeleteTodoTag(todoId int, tagId int) error
	
}
