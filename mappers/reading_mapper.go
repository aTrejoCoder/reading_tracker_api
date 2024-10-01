package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingMapper struct {
}

func (rm ReadingMapper) InsertDtoToEntity(readingInsertDTO dtos.ReadingInsertDTO) models.Reading {
	userId, _ := primitive.ObjectIDFromHex(readingInsertDTO.UserId)
	return models.Reading{
		UserId:          userId,
		ReadingType:     readingInsertDTO.ReadingType,
		ReadingsRecords: []models.ReadingRecord{},
		ReadingStatus:   readingInsertDTO.ReadingStatus,
		Notes:           readingInsertDTO.Notes,
		CreatedAt:       readingInsertDTO.CreatedAt,
		UpdatedAt:       readingInsertDTO.UpdatedAt,
	}
}

func (rm ReadingMapper) EntityToDTO(reading models.Reading) dtos.ReadingDTO {
	return dtos.ReadingDTO{
		Id:            reading.Id,
		ReadingType:   reading.ReadingType,
		ReadingStatus: reading.ReadingStatus,
		Notes:         reading.Notes,
		UpdatedAt:     reading.UpdatedAt,
	}
}

func (rm ReadingMapper) InsertDtoToUpdatedEntity(readingInsertDTO dtos.ReadingInsertDTO, currentReading models.Reading) models.Reading {
	currentReading.Notes = readingInsertDTO.Notes
	currentReading.ReadingStatus = readingInsertDTO.ReadingStatus
	currentReading.ReadingType = readingInsertDTO.ReadingType
	currentReading.UpdatedAt = time.Now().UTC()

	return currentReading
}
