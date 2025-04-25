package handler

import (
	"source-base-go/usecase/tag"
	"github.com/gin-gonic/gin"
	"source-base-go/api/middleware"
	"source-base-go/infrastructure/repository/util"
)

func MakeHandlers(app *gin.Engine, tagService tag.UseCase, verifier util.Verifier) {
	tagGroup := app.Group("/api/tag")
	{
		tagGroup.Use(middleware.AuthMiddleware(verifier))

		tagGroup.GET("/getAll", func(ctx *gin.Context) {
			getAllTag(ctx, tagService)
		})
		tagGroup.GET("/detail", func(ctx *gin.Context) {
			getTagDetail(ctx, tagService)
		})
		tagGroup.POST("/create", func(ctx *gin.Context) {
			createTag(ctx, tagService)
		})
		tagGroup.PUT("/update", func(ctx *gin.Context) {
			updateTag(ctx, tagService)
		})
		tagGroup.DELETE("/delete", func(ctx *gin.Context) {
			deleteTag(ctx, tagService)
		})
	}
}
