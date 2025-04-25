package handler

import (
	"fmt"
	"net/http"
	"strconv"

	tagPayload "source-base-go/api/payload/tag"
	tagPresenter "source-base-go/api/presenter/tag"
	"source-base-go/usecase/tag"

	"github.com/gin-gonic/gin"
)

func getAllTag(ctx *gin.Context, tagService tag.UseCase) {
	listTag, err := tagService.GetAllTag()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &tagPresenter.ListTagResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: "Success",
		Results: convertListTagEntityToPresenter(listTag),
	}

	ctx.JSON(http.StatusOK, response)
}

func getTagDetail(ctx *gin.Context, tagService tag.UseCase) {
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

	tag, err := tagService.GetTagDetail(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := &tagPresenter.TagResp{
		Status:  fmt.Sprint(http.StatusOK),
		Message: "Success",
		Results: convertTagEntityToPresenter(tag),
	}

	ctx.JSON(http.StatusOK, response)
}

func createTag(ctx *gin.Context, tagService tag.UseCase) {
	var payload tagPayload.TagPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := tagService.CreateTag(convertTagPayloadToEntity(&payload)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tag created successfully"})
}

func updateTag(ctx *gin.Context, tagService tag.UseCase) {
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

	var payload tagPayload.TagUpdatePayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := tagService.UpdateTag(id, convertTagUpdatePayloadToEntity(id, &payload)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tag updated successfully"})
}

func deleteTag(ctx *gin.Context, tagService tag.UseCase) {
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

	if err := tagService.DeleteTag(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
