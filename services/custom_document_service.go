package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomDocumentService interface {
	GetCustomDocumentsByUserId(userId primitive.ObjectID) ([]dtos.CustomDocumentDTO, error)
	GetCustomDocumentById(userId primitive.ObjectID, customDocumentId primitive.ObjectID) (*dtos.CustomDocumentDTO, error)
	CreateCustomDocument(customDocumentInsertDTO dtos.CustomDocumentInsertDTO, userId primitive.ObjectID) error
	UpdateCustomDocument(customDoctId primitive.ObjectID, userId primitive.ObjectID, customDocInsertDTO dtos.CustomDocumentInsertDTO) error
	DeleteCustomDocument(userId primitive.ObjectID, customDocumentId primitive.ObjectID) error
}

func NewCustomDocumentService(customDocumentRepository repository.CustomDocumentRepository) CustomDocumentService {
	return &CustomDocumentServiceImpl{
		customDocumentRepository: customDocumentRepository,
	}
}

type CustomDocumentServiceImpl struct {
	customDocumentRepository repository.CustomDocumentRepository
	CustomDocumentMapper     mappers.CustomDocumentMapper
}

func (ds CustomDocumentServiceImpl) GetCustomDocumentsByUserId(userId primitive.ObjectID) ([]dtos.CustomDocumentDTO, error) {
	customDocuments, err := ds.customDocumentRepository.GetByUserId(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	customDocumentsDTOs := []dtos.CustomDocumentDTO{}
	for _, customDocument := range customDocuments {
		customDocumnentDTO := ds.CustomDocumentMapper.EntityToDto(customDocument)
		customDocumentsDTOs = append(customDocumentsDTOs, customDocumnentDTO)
	}

	return customDocumentsDTOs, nil
}

func (ds CustomDocumentServiceImpl) GetCustomDocumentById(userId primitive.ObjectID, customDocumentId primitive.ObjectID) (*dtos.CustomDocumentDTO, error) {
	customDocument, err := ds.customDocumentRepository.GetById(context.TODO(), userId, customDocumentId)
	if err != nil {
		return nil, err
	}

	customDocumentDTO := ds.CustomDocumentMapper.EntityToDto(*customDocument)
	return &customDocumentDTO, nil
}

func (ds CustomDocumentServiceImpl) CreateCustomDocument(customDocumentInsertDTO dtos.CustomDocumentInsertDTO, userId primitive.ObjectID) error {
	customDocument := ds.CustomDocumentMapper.InsertDtoToEntity(customDocumentInsertDTO)

	err := ds.customDocumentRepository.Create(context.Background(), userId, customDocument)
	if err != nil {
		return err
	}

	return nil
}

func (ds CustomDocumentServiceImpl) UpdateCustomDocument(customDoctId primitive.ObjectID, userId primitive.ObjectID, customDocInsertDTO dtos.CustomDocumentInsertDTO) error {
	_, err := ds.customDocumentRepository.Update(context.Background(), userId, customDoctId, customDocInsertDTO)
	if err != nil {
		return err
	}
	return nil
}

func (ds CustomDocumentServiceImpl) DeleteCustomDocument(userId primitive.ObjectID, customDocumentId primitive.ObjectID) error {
	_, err := ds.customDocumentRepository.Delete(context.TODO(), userId, customDocumentId)
	if err != nil {
		return err
	}
	return nil
}
