package services

import (
	"context"
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingService interface {
	GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error)
	GetReadingsByUserId(userId primitive.ObjectID, page, limit int64, isAsc bool) ([]dtos.ReadingDTO, error)
	GetReadingsByUserAndType(userId primitive.ObjectID, sortValue, readingType string, page, limit int64, isAsc bool) ([]dtos.ReadingDTO, error)
	GetReadingsByUserAndStatus(userId primitive.ObjectID, readingStatus string, page, limit int64, isAsc bool) ([]dtos.ReadingDTO, error)

	CreateReading(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) error
	UpdateReading(readingId primitive.ObjectID, userId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error
	DeleteReading(readingId primitive.ObjectID, userId primitive.ObjectID) error
}

type readingServiceImpl struct {
	readingRepository repository.Repository[models.Reading]
	mangaRepository   repository.Repository[models.Manga]
	bookRepository    repository.Repository[models.Book]
	userRepository    repository.Repository[models.User]

	readingExtendRepository  repository.ReadingExtendRepository
	customDocumentRepository repository.CustomDocumentRepository

	readingMapper mappers.ReadingMapper
}

func NewReadingService(readingRepository repository.Repository[models.Reading],
	readingExtendRepository repository.ReadingExtendRepository,
	customDocumentRepository repository.CustomDocumentRepository,
	mangaRepository repository.Repository[models.Manga],
	bookRepository repository.Repository[models.Book],
	userRepository repository.Repository[models.User]) ReadingService {
	return &readingServiceImpl{
		readingRepository:        readingRepository,
		readingExtendRepository:  readingExtendRepository,
		mangaRepository:          mangaRepository,
		bookRepository:           bookRepository,
		userRepository:           userRepository,
		customDocumentRepository: customDocumentRepository,
	}
}

func (rs readingServiceImpl) GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error) {
	reading, err := rs.readingRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return nil, err
	}

	readingDTO := rs.readingMapper.EntityToDTO(*reading)
	return &readingDTO, nil
}

func (rs readingServiceImpl) GetReadingsByUserId(userId primitive.ObjectID, page, limit int64, isAsc bool) ([]dtos.ReadingDTO, error) {
	var sortOrder int
	if isAsc {
		sortOrder = 1 // asc
	} else {
		sortOrder = -1 // desc
	}

	readings, err := rs.readingExtendRepository.GetReadingsByUserId(context.Background(), userId, page, limit, sortOrder)

	if err != nil {
		return []dtos.ReadingDTO{}, err
	}

	readingDTOs := []dtos.ReadingDTO{}
	for _, reading := range readings {
		readingDTO := rs.readingMapper.EntityToDTO(reading)
		readingDTOs = append(readingDTOs, readingDTO)
	}

	return readingDTOs, nil
}

func (rs readingServiceImpl) GetReadingsByUserAndType(userId primitive.ObjectID, sortValue, readingType string, page, limit int64, isAsc bool) ([]dtos.ReadingDTO, error) {
	var sortOrder int
	if isAsc {
		sortOrder = 1 // asc
	} else {
		sortOrder = -1 // desc
	}

	readings, err := rs.readingExtendRepository.GetReadingsByUserIdAndReadingType(context.Background(), userId, readingType, sortValue, page, limit, sortOrder)

	if err != nil {
		return []dtos.ReadingDTO{}, err
	}

	readingDTOs := []dtos.ReadingDTO{}
	for _, reading := range readings {
		readingDTO := rs.readingMapper.EntityToDTO(reading)
		readingDTOs = append(readingDTOs, readingDTO)
	}

	return readingDTOs, nil
}

func (rs readingServiceImpl) GetReadingsByUserAndStatus(userId primitive.ObjectID, readingType string, page, limit int64, isAsc bool) ([]dtos.ReadingDTO, error) {
	var sortOrder int
	if isAsc {
		sortOrder = 1 // asc
	} else {
		sortOrder = -1 // desc
	}

	readings, err := rs.readingExtendRepository.GetReadingsByUserIdAndReadingStatus(context.Background(), userId, readingType, page, limit, sortOrder)

	if err != nil {
		return []dtos.ReadingDTO{}, err
	}

	readingDTOs := []dtos.ReadingDTO{}
	for _, reading := range readings {
		readingDTO := rs.readingMapper.EntityToDTO(reading)
		readingDTOs = append(readingDTOs, readingDTO)
	}

	return readingDTOs, nil
}

func (rs readingServiceImpl) CreateReading(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) error {
	documentName, err := rs.validateDocumentId(readingInsertDTO.ReadingType, readingInsertDTO.DocumentId, userId)
	if err != nil {
		return err
	}

	isReadingExists, err := rs.readingExtendRepository.IsReadingExistsForUser(context.TODO(), userId, readingInsertDTO.DocumentId)
	if err != nil {
		return err
	}

	if isReadingExists {
		return utils.ErrDuplicated
	}

	newReading := rs.readingMapper.InsertDtoToEntity(readingInsertDTO, userId, documentName)

	if _, err := rs.readingRepository.Create(context.TODO(), &newReading); err != nil {
		return err
	}

	return nil
}

func (rs readingServiceImpl) UpdateReading(readingId primitive.ObjectID, userId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error {
	_, err := rs.validateDocumentId(readingInsertDTO.ReadingType, readingInsertDTO.DocumentId, userId)
	if err != nil {
		return err
	}

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

func (rs readingServiceImpl) validateDocumentId(readingType string, documentId primitive.ObjectID, userId primitive.ObjectID) (string, error) {
	switch readingType {
	case "book":
		book, err := rs.bookRepository.GetByID(context.Background(), documentId)
		return book.Name, err
	case "manga":
		manga, err := rs.mangaRepository.GetByID(context.Background(), documentId)
		return manga.Title, err
	case "custom_document":
		customDocument, err := rs.customDocumentRepository.GetById(context.Background(), userId, documentId)
		return customDocument.Title, err
	default:
		return "", errors.New("invalid reading type")
	}
}
