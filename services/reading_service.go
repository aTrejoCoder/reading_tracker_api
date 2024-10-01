package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingService interface {
	GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error)
	CreateReading(readingInsertDTO dtos.ReadingInsertDTO) error
	UpdateReading(readingId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error
	DeleteReading(readingId primitive.ObjectID) error
}

type readingServiceImpl struct {
	commonRepository repository.Repository[models.Reading]
	readingMapper    mappers.ReadingMapper
}

func NewReadingService(commonRepository repository.Repository[models.Reading]) ReadingService {
	return &readingServiceImpl{
		commonRepository: commonRepository,
	}
}

func (bs readingServiceImpl) GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error) {
	reading, err := bs.commonRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return nil, err
	}

	ReadingDTO := bs.readingMapper.EntityToDTO(*reading)
	return &ReadingDTO, nil
}

func (bs readingServiceImpl) CreateReading(readingInsertDTO dtos.ReadingInsertDTO) error {
	newReading := bs.readingMapper.InsertDtoToEntity(readingInsertDTO)

	if _, err := bs.commonRepository.Create(context.TODO(), &newReading); err != nil {
		return err
	}

	return nil
}

func (bs readingServiceImpl) UpdateReading(readingId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error {
	currentReading, err := bs.commonRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return err
	}

	updatedReading := bs.readingMapper.InsertDtoToUpdatedEntity(readingInsertDTO, *currentReading)

	if _, err := bs.commonRepository.Update(context.TODO(), readingId, &updatedReading); err != nil {
		return err
	}

	return nil
}

func (bs readingServiceImpl) DeleteReading(readingId primitive.ObjectID) error {
	if _, err := bs.commonRepository.DeleteById(context.TODO(), readingId); err != nil {
		return err
	}
	return nil
}
