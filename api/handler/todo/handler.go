package handler

import (
	"source-base-go/usecase/todo"
	"github.com/gin-gonic/gin"
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/util"
)

func MakeHandlers(app *gin.Engine, todoService todo.UseCase, verifier util.Verifier) {
	todoGroup := app.Group("/api/todo")
	{
		todoGroup.Use(middleware.AuthMiddleware(verifier))

		todoGroup.GET("/getAll", func(ctx *gin.Context) {
			getAllTodo(ctx, todoService)
		})
		todoGroup.GET("/detail", func(ctx *gin.Context) {
			getTodoDetail(ctx, todoService)
		})
		todoGroup.POST("/create", func(ctx *gin.Context) {
			createTodo(ctx, todoService)
		})
		todoGroup.PUT("/update", func(ctx *gin.Context) {
			updateTodo(ctx, todoService)
		})
		todoGroup.DELETE("/delete", func(ctx *gin.Context) {
			deleteTodo(ctx, todoService)
		})
	}
}
