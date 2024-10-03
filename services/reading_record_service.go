package services

import (
	"context"
	"log"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingRecordService interface {
	CreateRecord(recordInsertDTO dtos.ReadingRecordInsertDTO, userId primitive.ObjectID) error
	GetRecordsByReadingId(readingId primitive.ObjectID) ([]dtos.ReadingRecordDTO, error)
	UpdateRecord(recordId primitive.ObjectID, recordUpdateDTO dtos.ReadingRecordInsertDTO) error
	DeleteRecord(readingId, userId primitive.ObjectID, recordId primitive.ObjectID) error

	GetRecordFromUser(readingId primitive.ObjectID, userId primitive.ObjectID) ([]dtos.ReadingRecordDTO, error)
	UpdateUserRecord(recordId primitive.ObjectID, userId primitive.ObjectID, recordUpdateDTO dtos.ReadingRecordInsertDTO) error
}

type readingRecordServiceImpl struct {
	readingMapper             mappers.ReadingMapper
	readingRepository         repository.Repository[models.Reading]
	readingExtendedRepository repository.ReadingExtendRepository
	UserService               UserService
}

func NewReadingRecordService(readingRepository repository.Repository[models.Reading],
	readingExtendedRepository repository.ReadingExtendRepository) ReadingRecordService {
	return &readingRecordServiceImpl{
		readingRepository:         readingRepository,
		readingExtendedRepository: readingExtendedRepository,
	}
}

func (rrs readingRecordServiceImpl) CreateRecord(recordInsertDTO dtos.ReadingRecordInsertDTO, userId primitive.ObjectID) error {
	log.Println("Starting CreateRecord function")

	newRecord := rrs.readingMapper.InsertRecordDtoToRecord(recordInsertDTO)
	log.Printf("New record created: %+v\n", newRecord)

	update := bson.M{
		"$push": bson.M{
			"reading_records": newRecord,
		},
	}

	log.Printf("Update query prepared: %+v\n", update)

	result, err := rrs.readingExtendedRepository.UpdateRecord(context.Background(), recordInsertDTO.ReadingId, userId, update)
	if err != nil {
		log.Printf("Error updating record: %v\n", err)
		return err
	}

	log.Printf("Record updated successfully: %+v\n", result)
	return nil
}

func (rrs readingRecordServiceImpl) GetRecordsByReadingId(readingId primitive.ObjectID) ([]dtos.ReadingRecordDTO, error) {
	reading, err := rrs.readingRepository.GetByID(context.Background(), readingId)
	if err != nil {
		return []dtos.ReadingRecordDTO{}, err
	}

	records := reading.ReadingsRecords

	var recordsDTO []dtos.ReadingRecordDTO
	for _, record := range records {
		recordDTO := rrs.readingMapper.RecordToDTO(record)
		recordsDTO = append(recordsDTO, recordDTO)
	}

	return recordsDTO, nil
}

func (rrs readingRecordServiceImpl) GetRecordFromUser(readingId primitive.ObjectID, userId primitive.ObjectID) ([]dtos.ReadingRecordDTO, error) {
	reading, err := rrs.readingRepository.GetByID(context.Background(), readingId)
	if err != nil {
		return []dtos.ReadingRecordDTO{}, err
	}

	if reading.UserId != userId {
		return []dtos.ReadingRecordDTO{}, utils.ErrForbidden
	}

	records := reading.ReadingsRecords

	var recordsDTO []dtos.ReadingRecordDTO
	for _, record := range records {
		recordDTO := rrs.readingMapper.RecordToDTO(record)
		recordsDTO = append(recordsDTO, recordDTO)
	}

	return recordsDTO, nil
}

func (rrs readingRecordServiceImpl) UpdateRecord(recordId primitive.ObjectID, recordUpdateDTO dtos.ReadingRecordInsertDTO) error {
	update := bson.M{
		"$set": bson.M{
			"reading_records.$[elem].notes":    recordUpdateDTO.Notes,
			"reading_records.$[elem].progress": recordUpdateDTO.Progress,
		},
	}

	arrayFilter := bson.A{
		bson.M{"elem._id": recordId},
	}

	_, err := rrs.readingExtendedRepository.UpdateWithFilter(
		context.Background(),
		recordUpdateDTO.ReadingId,
		update,
		arrayFilter,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rrs readingRecordServiceImpl) UpdateUserRecord(recordId primitive.ObjectID, userId primitive.ObjectID, recordUpdateDTO dtos.ReadingRecordInsertDTO) error {
	_, err := rrs.readingExtendedRepository.UpdateRecordProperties(
		context.Background(),
		recordUpdateDTO.ReadingId,
		recordId,
		recordUpdateDTO,
	)
	if err != nil {
		return err
	}

	return nil
}

func (rrs readingRecordServiceImpl) DeleteRecord(readingId, userId primitive.ObjectID, recordId primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"reading_records": bson.M{"_id": recordId},
		},
	}

	_, err := rrs.readingExtendedRepository.UpdateRecord(context.Background(), readingId, userId, update)
	if err != nil {
		return err
	}

	return nil
}
