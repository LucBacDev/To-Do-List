package handler

import (
	"github.com/gin-gonic/gin"
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/util"
	"source-base-go/usecase/todotag"

)

func MakeHandlers(app *gin.Engine, todoTagService todotag.UseCase, verifier util.Verifier) {
	todoTagGroup := app.Group("/api/todo-tag")
	{
		todoTagGroup.Use(middleware.AuthMiddleware(verifier))

		todoTagGroup.POST("/create", func(ctx *gin.Context) {
			createTodoTag(ctx, todoTagService)
		})
		todoTagGroup.DELETE("/delete", func(ctx *gin.Context) {
			deleteTodoTag(ctx, todoTagService)
		})
	}
}
