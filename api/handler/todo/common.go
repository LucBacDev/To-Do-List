package handler

import (
	todoPayload "source-base-go/api/payload/todo"
	todoPresenter "source-base-go/api/presenter/todo"
	"source-base-go/entity"
	"time"

	"github.com/gin-gonic/gin"
)

func convertListTodoEntityToPresenter(ctx *gin.Context, listData []*entity.Todo) []*todoPresenter.Todo {

	listTodoPresenter := []*todoPresenter.Todo{}
	userId := ctx.GetInt("user_id")
	for _, todo := range listData {
		todoPresenter := &todoPresenter.Todo{
			Id:          todo.Id,
			UserId:      userId,
			Title:       todo.Title,
			Description: todo.Description,
			IsDone:      todo.IsDone,
			Priority:    string(todo.Priority),
			DueDate:     formatTime(todo.DueDate),
			CreatedAt:   todo.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   todo.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		listTodoPresenter = append(listTodoPresenter, todoPresenter)
	}

	return listTodoPresenter
}
func convertTodoEntityToPresenter(ctx *gin.Context, id int, data *entity.Todo) *todoPresenter.Todo {
	userId := ctx.GetInt("user_id")
	return &todoPresenter.Todo{
		Id:          id,
		UserId:      userId,
		Title:       data.Title,
		Description: data.Description,
		IsDone:      data.IsDone,
		Priority:    string(data.Priority),
		DueDate:     formatTime(data.DueDate),
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func convertTodoPayloadToEntity(ctx *gin.Context, data *todoPayload.TodoPayload) *entity.Todo {
	userId := ctx.GetInt("user_id")

	return &entity.Todo{
		UserId:      userId,
		Title:       data.Title,
		Description: data.Description,
		Priority:    entity.Priority(data.Priority),
		DueDate:     parseTime(data.DueDate),
	}
}

func convertTodoUpdatePayloadToEntity(ctx *gin.Context, id int, data *todoPayload.TodoUpdatePayload) *entity.Todo {
	userId := ctx.GetInt("user_id")

	return &entity.Todo{
		Id:          id,
		UserId:      userId,
		Title:       data.Title,
		Description: data.Description,
		IsDone:      data.IsDone,
		Priority:    entity.Priority(data.Priority),
		DueDate:     parseTime(data.DueDate),
	}
}

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func parseTime(str string) *time.Time {
	if str == "" {
		return nil
	}
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return nil
	}
	return &t
}
