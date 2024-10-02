package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingMapper struct {
}

func (rm ReadingMapper) InsertDtoToEntity(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) models.Reading {
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
	records := reading.ReadingsRecords

	recordsDTOs := []dtos.ReadingRecordDTO{}

	for _, record := range records {
		recordDTO := rm.RecordToDTO(record)
		recordsDTOs = append(recordsDTOs, recordDTO)
	}

	return dtos.ReadingDTO{
		Id:            reading.Id,
		ReadingType:   reading.ReadingType,
		UserId:        reading.UserId.Hex(),
		ReadingStatus: reading.ReadingStatus,
		Notes:         reading.Notes,
		RecordsDTOs:   recordsDTOs,
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

func (rm ReadingMapper) InsertRecordDtoToRecord(recordDTO dtos.ReadingRecordInsertDTO) models.ReadingRecord {
	return models.ReadingRecord{
		Id:         primitive.NewObjectID(),
		Notes:      recordDTO.Notes,
		Progress:   recordDTO.Progress,
		RecordDate: time.Now().UTC(),
	}
}

func (rm ReadingMapper) RecordToDTO(recordDTO models.ReadingRecord) dtos.ReadingRecordDTO {
	return dtos.ReadingRecordDTO{
		Id:         recordDTO.Id,
		Notes:      recordDTO.Notes,
		Progress:   recordDTO.Progress,
		RecordDate: recordDTO.RecordDate,
	}
}
