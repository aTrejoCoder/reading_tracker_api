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
	GetReadingsByUserId(readingId primitive.ObjectID) ([]dtos.ReadingDTO, error)

	CreateReading(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) error
	UpdateReading(readingId primitive.ObjectID, userId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error
	DeleteReading(readingId primitive.ObjectID, userId primitive.ObjectID) error
}

type readingServiceImpl struct {
	readingRepository repository.Repository[models.Reading]
	mangaRepository   repository.Repository[models.Manga]
	bookRepository    repository.Repository[models.Book]
	userRepository    repository.Repository[models.User]

	readingExtendRepository repository.ReadingExtendRepository
	readingMapper           mappers.ReadingMapper
	bookMapper              mappers.BookMapper
	mangaMapper             mappers.MangaMapper
}

func NewReadingService(readingRepository repository.Repository[models.Reading],
	readingExtendRepository repository.ReadingExtendRepository,
	mangaRepository repository.Repository[models.Manga],
	bookRepository repository.Repository[models.Book],
	userRepository repository.Repository[models.User]) ReadingService {
	return &readingServiceImpl{
		readingRepository:       readingRepository,
		readingExtendRepository: readingExtendRepository,
		mangaRepository:         mangaRepository,
		bookRepository:          bookRepository,
		userRepository:          userRepository,
	}
}

func (rs readingServiceImpl) GetReadingById(readingId primitive.ObjectID) (*dtos.ReadingDTO, error) {
	reading, err := rs.readingRepository.GetByID(context.TODO(), readingId)
	if err != nil {
		return nil, err
	}

	readingDTO := rs.readingMapper.EntityToDTO(*reading)
	documentDTO := rs.fetchDocumentData(reading.ReadingType, reading.DocumentId)
	readingDTO.DocumentData = documentDTO

	return &readingDTO, nil
}

func (rs readingServiceImpl) GetReadingsByUserId(userId primitive.ObjectID) ([]dtos.ReadingDTO, error) {
	readings, err := rs.readingExtendRepository.GetReadingsByUserId(context.Background(), userId)

	if err != nil {
		return []dtos.ReadingDTO{}, err
	}

	var readingDTOs []dtos.ReadingDTO
	for _, reading := range readings {

		readingDTO := rs.readingMapper.EntityToDTO(reading)
		documentDTO := rs.fetchDocumentData(reading.ReadingType, reading.DocumentId)

		readingDTO.DocumentData = documentDTO

		readingDTOs = append(readingDTOs, readingDTO)
	}

	return readingDTOs, nil
}

func (rs readingServiceImpl) CreateReading(readingInsertDTO dtos.ReadingInsertDTO, userId primitive.ObjectID) error {
	err := rs.validateDocumentId(readingInsertDTO.ReadingType, readingInsertDTO.DocumentId)
	if err != nil {
		return err
	}

	newReading := rs.readingMapper.InsertDtoToEntity(readingInsertDTO, userId)

	if _, err := rs.readingRepository.Create(context.TODO(), &newReading); err != nil {
		return err
	}

	return nil
}

func (rs readingServiceImpl) UpdateReading(readingId primitive.ObjectID, userId primitive.ObjectID, readingInsertDTO dtos.ReadingInsertDTO) error {
	err := rs.validateDocumentId(readingInsertDTO.ReadingType, readingInsertDTO.DocumentId)
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

func (rs readingServiceImpl) validateDocumentId(readingType string, documentId primitive.ObjectID) error {
	switch readingType {
	case "book":
		_, err := rs.bookRepository.GetByID(context.Background(), documentId)
		return err
	case "manga":
		_, err := rs.mangaRepository.GetByID(context.Background(), documentId)
		return err
	case "custom_document":
		return nil
	default:
		return errors.New("invalid reading type")
	}
}

func (rs readingServiceImpl) fetchDocumentData(readingType string, documentId primitive.ObjectID) interface{} {
	switch readingType {
	case "book":
		book, err := rs.bookRepository.GetByID(context.Background(), documentId)
		if err != nil || book == nil {
			return nil
		}
		return rs.bookMapper.EntityToDTO(*book)
	case "manga":
		manga, err := rs.mangaRepository.GetByID(context.Background(), documentId)
		if err != nil || manga == nil {
			return nil
		}
		return rs.mangaMapper.EntityToDTO(*manga)
	case "custom_document":
		return nil
	default:
		return nil
	}
}
