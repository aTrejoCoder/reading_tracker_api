package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MangaService interface {
	GetMangaId(MangaId primitive.ObjectID) (*dtos.MangaDTO, error)
	CreateManga(MangaInsertDTO dtos.MangaInsertDTO) error
	UpdateManga(MangaId primitive.ObjectID, MangaInsertDTO dtos.MangaInsertDTO) error
	DeleteManga(MangaId primitive.ObjectID) error
}

type MangaServiceImpl struct {
	MangaMapper      mappers.MangaMapper
	commonRepository repository.Repository[models.Manga]
}

func NewMangaService(commonRepository repository.Repository[models.Manga]) MangaService {
	return &MangaServiceImpl{
		commonRepository: commonRepository,
	}

}

func (bs MangaServiceImpl) GetMangaId(MangaId primitive.ObjectID) (*dtos.MangaDTO, error) {
	manga, err := bs.commonRepository.GetByID(context.TODO(), MangaId)
	if err != nil {
		return nil, err
	}

	mangaDTO := bs.MangaMapper.EntityToDTO(*manga)
	return &mangaDTO, nil
}

func (bs MangaServiceImpl) CreateManga(MangaInsertDTO dtos.MangaInsertDTO) error {
	newManga := bs.MangaMapper.InsertDtoToEntity(MangaInsertDTO)

	if _, err := bs.commonRepository.Create(context.TODO(), &newManga); err != nil {
		return err
	}

	return nil
}

func (bs MangaServiceImpl) UpdateManga(MangaId primitive.ObjectID, MangaInsertDTO dtos.MangaInsertDTO) error {
	mangaUpdated := bs.MangaMapper.InsertDtoToEntity(MangaInsertDTO)

	if _, err := bs.commonRepository.Update(context.TODO(), MangaId, &mangaUpdated); err != nil {
		return err
	}

	return nil
}

func (bs MangaServiceImpl) DeleteManga(MangaId primitive.ObjectID) error {
	if _, err := bs.commonRepository.Delete(context.TODO(), MangaId); err != nil {
		return err
	}
	return nil
}
