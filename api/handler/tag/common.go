package handler

import (
	tagPayload "source-base-go/api/payload/tag"
	tagPresenter "source-base-go/api/presenter/tag"
	"source-base-go/entity"
)

func convertListTagEntityToPresenter(listData []*entity.Tag) []*tagPresenter.Tag {
	listTagPresenter := []*tagPresenter.Tag{}
	for _, tag := range listData {
		tagPresenter := &tagPresenter.Tag{
			Id:        tag.Id,
			Name:      tag.Name,
			Color:     tag.Color,
			CreatedAt: tag.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: tag.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		listTagPresenter = append(listTagPresenter, tagPresenter)
	}
	return listTagPresenter
}

func convertTagEntityToPresenter(data *entity.Tag) *tagPresenter.Tag {
	return &tagPresenter.Tag{
		Id:        data.Id,
		Name:      data.Name,
		Color:     data.Color,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func convertTagPayloadToEntity(data *tagPayload.TagPayload) *entity.Tag {
	return &entity.Tag{
		Name:  data.Name,
		Color: data.Color,
	}
}

func convertTagUpdatePayloadToEntity(id int, data *tagPayload.TagUpdatePayload) *entity.Tag {
	return &entity.Tag{
		Id:    id,
		Name:  data.Name,
		Color: data.Color,
	}
}
