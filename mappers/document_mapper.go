package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentMapper struct {
}

func (dm DocumentMapper) InsertDtoToEntity(documentInsertDTO dtos.DocumentInsertDTO) models.Document {
	now := time.Now().UTC()
	return models.Document{
		Title:       documentInsertDTO.Title,
		Author:      documentInsertDTO.Author,
		Description: documentInsertDTO.Description,
		FileURL:     documentInsertDTO.FileURL,
		FileType:    documentInsertDTO.FileType,
		Tags:        documentInsertDTO.Tags,
		Version:     documentInsertDTO.Version,
		Status:      documentInsertDTO.Status,
		ReadingList: []primitive.ObjectID{},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (dm DocumentMapper) EntityToDto(document models.Document) dtos.DocumentDTO {
	return dtos.DocumentDTO{
		Id:          document.Id,
		Title:       document.Title,
		Author:      document.Author,
		Description: document.Description,
		FileURL:     document.FileURL,
		FileType:    document.FileType,
		Tags:        document.Tags,
		Version:     document.Version,
		Status:      document.Status,
	}
}

func (dm DocumentMapper) InsertDtoToUpdatedEntity(insertDTO dtos.DocumentInsertDTO, currentDocument models.Document) models.Document {
	currentDocument.Title = insertDTO.Title
	currentDocument.Author = insertDTO.Author
	currentDocument.Description = insertDTO.Description
	currentDocument.FileURL = insertDTO.FileURL
	currentDocument.FileType = insertDTO.FileType
	currentDocument.Tags = insertDTO.Tags
	currentDocument.Version = insertDTO.Version
	currentDocument.Status = insertDTO.Status

	return currentDocument
}
