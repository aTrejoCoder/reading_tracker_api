package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentService interface {
	GetDocumentById(documentId primitive.ObjectID) (*dtos.DocumentDTO, error)
	CreateDocument(documentInsertDTO dtos.DocumentInsertDTO) error
	UpdateDocument(documentId primitive.ObjectID, documentInsertDTO dtos.DocumentInsertDTO) error
	DeleteDocument(documentId primitive.ObjectID) error
}

func NewDocumentService(commonRepository repository.Repository[models.Document]) DocumentService {
	return &documentServiceImpl{
		commonRepository: commonRepository,
	}
}

type documentServiceImpl struct {
	commonRepository repository.Repository[models.Document]
	documentMapper   mappers.DocumentMapper
}

func (ds documentServiceImpl) GetDocumentById(documentId primitive.ObjectID) (*dtos.DocumentDTO, error) {
	document, err := ds.commonRepository.GetByID(context.TODO(), documentId)
	if err != nil {
		return nil, err
	}

	documnentDTO := ds.documentMapper.EntityToDto(*document)
	return &documnentDTO, nil
}

func (ds documentServiceImpl) CreateDocument(documentInsertDTO dtos.DocumentInsertDTO) error {
	document := ds.documentMapper.InsertDtoToEntity(documentInsertDTO)
	_, err := ds.commonRepository.Create(context.TODO(), &document)
	if err != nil {
		return err
	}
	return nil
}

func (ds documentServiceImpl) UpdateDocument(documentId primitive.ObjectID, documentInsertDTO dtos.DocumentInsertDTO) error {
	currentDocument, err := ds.commonRepository.GetByID(context.TODO(), documentId)
	if err != nil {
		return err
	}

	updatedEntity := ds.documentMapper.InsertDtoToUpdatedEntity(documentInsertDTO, *currentDocument)

	_, err = ds.commonRepository.Update(context.TODO(), documentId, updatedEntity)
	if err != nil {
		return err
	}
	return nil
}

func (ds documentServiceImpl) DeleteDocument(documentId primitive.ObjectID) error {
	_, err := ds.commonRepository.DeleteById(context.TODO(), documentId)
	if err != nil {
		return err
	}
	return nil
}
