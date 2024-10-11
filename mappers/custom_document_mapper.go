package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomDocumentMapper struct {
}

func (dm CustomDocumentMapper) InsertDtoToEntity(customDocumentInsertDTO dtos.CustomDocumentInsertDTO) models.CustomDocument {
	now := time.Now().UTC()
	return models.CustomDocument{
		Id:          primitive.NewObjectID(),
		Title:       customDocumentInsertDTO.Title,
		Author:      customDocumentInsertDTO.Author,
		Description: customDocumentInsertDTO.Description,
		FileURL:     customDocumentInsertDTO.FileURL,
		Tags:        customDocumentInsertDTO.Tags,
		URL:         customDocumentInsertDTO.URL,
		Content:     customDocumentInsertDTO.Content,
		Category:    customDocumentInsertDTO.Category,
		Version:     customDocumentInsertDTO.Version,
		Status:      customDocumentInsertDTO.Status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

}

func (dm CustomDocumentMapper) EntityToDto(customDocument models.CustomDocument) dtos.CustomDocumentDTO {
	return dtos.CustomDocumentDTO{
		Id:          customDocument.Id,
		Title:       customDocument.Title,
		Author:      customDocument.Author,
		Description: customDocument.Description,
		FileURL:     customDocument.FileURL,
		Tags:        customDocument.Tags,
		URL:         customDocument.URL,
		Content:     customDocument.Content,
		Category:    customDocument.Category,
		Version:     customDocument.Version,
		Status:      customDocument.Status,
	}
}

func (dm CustomDocumentMapper) InsertDtoToUpdateDocument(customDocumentInsertDTO dtos.CustomDocumentInsertDTO) primitive.M {
	return bson.M{
		"title":       customDocumentInsertDTO.Title,
		"author":      customDocumentInsertDTO.Author,
		"description": customDocumentInsertDTO.Description,
		"file_url":    customDocumentInsertDTO.FileURL,
		"tags":        customDocumentInsertDTO.Tags,
		"url":         customDocumentInsertDTO.URL,
		"category":    customDocumentInsertDTO.Category,
		"version":     customDocumentInsertDTO.Version,
		"status":      customDocumentInsertDTO.Status,
		"updated_at":  time.Now().UTC(),
	}
}
