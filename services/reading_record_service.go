package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingRecordService interface {
	CreateRecord(recordInsertDTO dtos.ReadingRecordInsertDTO) error
	GetRecordsByReadingId(readingId primitive.ObjectID) ([]dtos.ReadingRecordDTO, error)
	UpdateRecord(recordId primitive.ObjectID, recordUpdateDTO dtos.ReadingRecordInsertDTO) error
	DeleteRecord(readingId, recordId primitive.ObjectID) error
}

type readingRecordServiceImpl struct {
	readingMapper             mappers.ReadingMapper
	readingRepository         repository.Repository[models.Reading]
	readingExtendedRepository repository.ReadingExtendRepository
}

func NewReadingRecordService(readingRepository repository.Repository[models.Reading],
	readingExtendedRepository repository.ReadingExtendRepository) ReadingRecordService {
	return &readingRecordServiceImpl{
		readingRepository: readingRepository,
	}
}

func (rrs readingRecordServiceImpl) CreateRecord(recordInsertDTO dtos.ReadingRecordInsertDTO) error {
	newRecord := rrs.readingMapper.InsertRecordDtoToRecord(recordInsertDTO)
	update := bson.M{
		"$push": bson.M{
			"reading_records": newRecord,
		},
	}

	_, err := rrs.readingExtendedRepository.UpdateRecord(context.Background(), recordInsertDTO.ReadingId, update)
	if err != nil {
		return err
	}

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

func (rrs readingRecordServiceImpl) DeleteRecord(readingId, recordId primitive.ObjectID) error {
	update := bson.M{
		"$pull": bson.M{
			"reading_records": bson.M{"_id": recordId},
		},
	}

	_, err := rrs.readingExtendedRepository.UpdateRecord(context.Background(), readingId, update)
	if err != nil {
		return err
	}

	return nil
}
