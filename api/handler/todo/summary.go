package handler

import (
	"fmt"
	"net/http"
	"strconv"

	todoPayload "source-base-go/api/payload/todo"
	todoPresenter "source-base-go/api/presenter/todo"
	"source-base-go/usecase/todo"

	"github.com/gin-gonic/gin"
)

func getAllTodo(ctx *gin.Context, todoService todo.UseCase) {
	listTodo, err := todoService.GetAllTodo()
	fmt.Println("List Todo:")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &todoPresenter.ListTodoResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: "Success",
		Results: convertListTodoEntityToPresenter(ctx, listTodo),
	}

	ctx.JSON(http.StatusOK, response)
}

func getTodoDetail(ctx *gin.Context, todoService todo.UseCase) {
	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
		return
	}

	todo, err := todoService.GetTodoDetail((id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := &todoPresenter.TodoResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: "Success",
		Results: convertTodoEntityToPresenter(ctx, id, todo),
	}

	ctx.JSON(http.StatusOK, response)

}

func createTodo(ctx *gin.Context, todoService todo.UseCase) {
	var payload todoPayload.TodoPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = todoService.CreateTodo(convertTodoPayloadToEntity(ctx, &payload))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
}

func updateTodo(ctx *gin.Context, todoService todo.UseCase) {
	var payload todoPayload.TodoUpdatePayload

	idStr := ctx.DefaultQuery("id", "") 
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
		return
	}

	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Println("payload:", payload)

	err = todoService.UpdateTodo(id, convertTodoUpdatePayloadToEntity(ctx, id, &payload))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func deleteTodo(ctx *gin.Context, todoService todo.UseCase) {
	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
		return
	}

	err = todoService.DeleteTodo((id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
