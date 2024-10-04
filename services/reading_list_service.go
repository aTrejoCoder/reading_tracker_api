package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingListService interface {
	AddReadingToList(userId primitive.ObjectID, moveReadingsToListDTO dtos.UpdateReadingsToListDTO) error
	RemoveReadingToList(userId primitive.ObjectID, moveReadingsToListDTO dtos.UpdateReadingsToListDTO) error

	GetReadingListsByUserId(userId primitive.ObjectID) ([]dtos.ReadingListDTO, error)
	GetReadingListById(listId primitive.ObjectID, userId primitive.ObjectID) (*dtos.ReadingListDTO, error)
	CreateReadingList(userId primitive.ObjectID, ReadingListInsertDTO dtos.ReadingListInsertDTO) error
	UpdateReadingList(readingListId primitive.ObjectID, userId primitive.ObjectID, insertDTO dtos.ReadingListInsertDTO) error
	DeleteReadingList(userId primitive.ObjectID, readingListId primitive.ObjectID) error
}

type readingListServiceImpl struct {
	readingListRepository repository.ReadingListRepository
	readingListMapper     mappers.ReadingListMapper
}

func NewReadingListService(readingListRepository repository.ReadingListRepository) ReadingListService {
	return &readingListServiceImpl{
		readingListRepository: readingListRepository,
	}
}

func (rls readingListServiceImpl) GetReadingListsByUserId(userId primitive.ObjectID) ([]dtos.ReadingListDTO, error) {
	readingLists, err := rls.readingListRepository.GetByUserId(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	readingListDTOs := []dtos.ReadingListDTO{}
	for _, readingList := range readingLists {
		readingListDTO := rls.readingListMapper.EntityToDTO(readingList)
		readingListDTOs = append(readingListDTOs, readingListDTO)
	}

	return readingListDTOs, nil
}

func (rls readingListServiceImpl) GetReadingListById(listId primitive.ObjectID, userId primitive.ObjectID) (*dtos.ReadingListDTO, error) {
	readingLists, err := rls.readingListRepository.GetByUserId(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	var readingListDTO *dtos.ReadingListDTO
	for _, readingList := range readingLists {
		if readingList.Id == listId {
			readingListDTO = &dtos.ReadingListDTO{}
			*readingListDTO = rls.readingListMapper.EntityToDTO(readingList)
			break
		}
	}

	if readingListDTO == nil {
		return nil, utils.ErrNotFound
	}

	return readingListDTO, nil
}

func (rls readingListServiceImpl) CreateReadingList(userId primitive.ObjectID, ReadingListInsertDTO dtos.ReadingListInsertDTO) error {
	readingList := rls.readingListMapper.InsertDtoToEntity(ReadingListInsertDTO)
	err := rls.readingListRepository.CreateReadingList(context.TODO(), userId, readingList)
	if err != nil {
		return err
	}
	return nil
}

func (rls readingListServiceImpl) UpdateReadingList(readingListId primitive.ObjectID, userId primitive.ObjectID, insertDTO dtos.ReadingListInsertDTO) error {
	if _, err := rls.readingListRepository.UpdateReadingList(context.Background(), userId, readingListId, insertDTO); err != nil {
		return err
	}
	return nil
}

func (rls readingListServiceImpl) DeleteReadingList(userId primitive.ObjectID, readingListId primitive.ObjectID) error {
	_, err := rls.readingListRepository.DeleteReadingList(context.TODO(), userId, readingListId)
	if err != nil {
		return err
	}
	return nil
}

func (rls readingListServiceImpl) AddReadingToList(userId primitive.ObjectID, updateReadingsToListDTO dtos.UpdateReadingsToListDTO) error {
	readingsIds := updateReadingsToListDTO.ReadingIds
	readingListId := updateReadingsToListDTO.ReadingListId
	if _, err := rls.readingListRepository.AddReadingsToList(context.Background(), userId, readingListId, readingsIds); err != nil {
		return err
	}

	return nil
}

func (rls readingListServiceImpl) RemoveReadingToList(userId primitive.ObjectID, updateReadingsToListDTO dtos.UpdateReadingsToListDTO) error {
	readingsIds := updateReadingsToListDTO.ReadingIds
	readingListId := updateReadingsToListDTO.ReadingListId
	if _, err := rls.readingListRepository.RemoveReadingsFromList(context.Background(), userId, readingListId, readingsIds); err != nil {
		return err
	}

	return nil
}
