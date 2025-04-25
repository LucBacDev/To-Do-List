package handler

import (
	"net/http"
	"strconv"

	"source-base-go/usecase/todotag"

	"github.com/gin-gonic/gin"
)

func createTodoTag(ctx *gin.Context, todoTagService todotag.UseCase) {
	todoIDStr := ctx.DefaultQuery("todo_id", "")
	tagIDStr := ctx.DefaultQuery("tag_id", "")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo_id"})
		return
	}

	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag_id"})
		return
	}
	err = todoTagService.CreateTodoTag(todoID, tagID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
}


func deleteTodoTag(ctx *gin.Context, todoTagService todotag.UseCase) {
	todoIDStr := ctx.DefaultQuery("todo_id", "")
	tagIDStr := ctx.DefaultQuery("tag_id", "")
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo_id"})
		return
	}

	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag_id"})
		return
	}
	err = todoTagService.DeleteTodoTag(todoID, tagID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo Tag deleted successfully"})
}
