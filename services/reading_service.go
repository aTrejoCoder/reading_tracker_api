package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingService interface {
	GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error)
	GetReadingsByUserId(readingId primitive.ObjectID) ([]dtos.ReadingDTO, error)
	CreateReading(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) error
	UpdateReading(readingId primitive.ObjectID, userId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error
	DeleteReading(readingId primitive.ObjectID, userId primitive.ObjectID) error
}

type readingServiceImpl struct {
	readingRepository       repository.Repository[models.Reading]
	readingExtendRepository repository.ReadingExtendRepository
	userRepository          repository.UserExtendRepository
	readingMapper           mappers.ReadingMapper
}

func NewReadingService(readingRepository repository.Repository[models.Reading],
	readingExtendRepository repository.ReadingExtendRepository,
	userRepository repository.UserExtendRepository) ReadingService {
	return &readingServiceImpl{
		readingRepository:       readingRepository,
		readingExtendRepository: readingExtendRepository,
		userRepository:          userRepository,
	}
}

func (rs readingServiceImpl) GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error) {
	reading, err := rs.readingRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return nil, err
	}

	ReadingDTO := rs.readingMapper.EntityToDTO(*reading)
	return &ReadingDTO, nil
}

func (rs readingServiceImpl) GetReadingsByUserId(userId primitive.ObjectID) ([]dtos.ReadingDTO, error) {
	readings, err := rs.readingExtendRepository.GetReadingsByUserId(context.Background(), userId)

	if err != nil {
		return []dtos.ReadingDTO{}, err
	}

	var readingDTOs []dtos.ReadingDTO
	for _, reading := range readings {
		readingDTO := rs.readingMapper.EntityToDTO(reading)
		readingDTOs = append(readingDTOs, readingDTO)
	}

	return readingDTOs, nil
}

func (rs readingServiceImpl) CreateReading(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) error {
	newReading := rs.readingMapper.InsertDtoToEntity(readingInsertDTO, userId)

	if _, err := rs.readingRepository.Create(context.TODO(), &newReading); err != nil {
		return err
	}

	return nil
}

func (rs readingServiceImpl) UpdateReading(readingId primitive.ObjectID, userId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error {
	currentReading, err := rs.readingRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return err
	}

	if currentReading.UserId != userId {
		return utils.ErrForbidden
	}

	updatedReading := rs.readingMapper.InsertDtoToUpdatedEntity(readingInsertDTO, *currentReading)

	if _, err := rs.readingRepository.Update(context.TODO(), readingId, &updatedReading); err != nil {
		return err
	}

	return nil
}

func (rs readingServiceImpl) DeleteReading(readingId primitive.ObjectID, userId primitive.ObjectID) error {
	reading, err := rs.readingRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return err
	}

	if reading.UserId != userId {
		return utils.ErrForbidden
	}

	if _, err := rs.readingRepository.DeleteById(context.TODO(), readingId); err != nil {
		return err
	}
	return nil
}
