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

type MangaService interface {
	GetMangaById(MangaId primitive.ObjectID) (*dtos.MangaDTO, error)
	GetMangaByNamePattern(mangaName string, page int64, limit int64) ([]dtos.MangaDTO, error)
	GetMangaByAuthor(author string) ([]dtos.MangaDTO, error)
	GetMangaByGenre(genre string, page int64, limit int64) ([]dtos.MangaDTO, error)
	GetMangaByDemography(genre string, page int64, limit int64) ([]dtos.MangaDTO, error)
	GetAllMangaSortedPaginated(page int64, limit int64) ([]dtos.MangaDTO, error)

	CreateManga(MangaInsertDTO dtos.MangaInsertDTO) error
	UpdateManga(MangaId primitive.ObjectID, MangaInsertDTO dtos.MangaInsertDTO) error
	DeleteManga(MangaId primitive.ObjectID) error
}

type MangaServiceImpl struct {
	mangaMapper     mappers.MangaMapper
	mangaRepository repository.Repository[models.Manga]
}

func NewMangaService(mangaRepository repository.Repository[models.Manga]) MangaService {
	return &MangaServiceImpl{
		mangaRepository: mangaRepository,
	}

}

func (bs MangaServiceImpl) GetMangaById(MangaId primitive.ObjectID) (*dtos.MangaDTO, error) {
	manga, err := bs.mangaRepository.GetByID(context.TODO(), MangaId)
	if err != nil {
		return nil, err
	}

	mangaDTO := bs.mangaMapper.EntityToDTO(*manga)
	return &mangaDTO, nil
}

func (bs MangaServiceImpl) GetMangaByNamePattern(mangaName string, page int64, limit int64) ([]dtos.MangaDTO, error) {
	filter := bson.M{"title": bson.M{"$regex": "^" + mangaName, "$options": "i"}}

	manga, err := bs.mangaRepository.GetManyByFilterPaginated(context.TODO(), filter, page, limit)
	if err != nil {
		return nil, err
	}

	var mangaDTOs []dtos.MangaDTO
	for _, manga := range manga {
		mangaDTOs = append(mangaDTOs, bs.mangaMapper.EntityToDTO(manga))
	}
	return mangaDTOs, nil
}

func (bs MangaServiceImpl) GetMangaByAuthor(author string) ([]dtos.MangaDTO, error) {
	filter := bson.M{"author": bson.M{"$regex": author, "$options": "i"}}

	mangas, err := bs.mangaRepository.GetManyByFilter(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var mangaDTOs []dtos.MangaDTO
	for _, manga := range mangas {
		mangaDTOs = append(mangaDTOs, bs.mangaMapper.EntityToDTO(manga))
	}
	return mangaDTOs, nil
}

func (bs MangaServiceImpl) GetMangaByGenre(genre string, page int64, limit int64) ([]dtos.MangaDTO, error) {
	filter := bson.M{"genres": bson.M{"$regex": genre, "$options": "i"}}

	Manga, err := bs.mangaRepository.GetManyByFilterPaginated(context.TODO(), filter, page, limit)
	if err != nil {
		return nil, err
	}

	var MangaDTOs []dtos.MangaDTO
	for _, Manga := range Manga {
		MangaDTOs = append(MangaDTOs, bs.mangaMapper.EntityToDTO(Manga))
	}
	return MangaDTOs, nil
}

func (bs MangaServiceImpl) GetMangaByDemography(genre string, page int64, limit int64) ([]dtos.MangaDTO, error) {
	filter := bson.M{"demograhpy": bson.M{"$regex": genre, "$options": "i"}}

	Manga, err := bs.mangaRepository.GetManyByFilterPaginated(context.TODO(), filter, page, limit)
	if err != nil {
		return nil, err
	}

	var MangaDTOs []dtos.MangaDTO
	for _, Manga := range Manga {
		MangaDTOs = append(MangaDTOs, bs.mangaMapper.EntityToDTO(Manga))
	}
	return MangaDTOs, nil
}

func (bs MangaServiceImpl) GetAllMangaSortedPaginated(page int64, limit int64) ([]dtos.MangaDTO, error) {
	sortFields := bson.D{
		{Key: "name", Value: 1},
		{Key: "author", Value: 1},
	}

	Manga, err := bs.mangaRepository.GetAllSortedPaginated(context.TODO(), sortFields, page, limit)
	if err != nil {
		return nil, err
	}

	var MangaDTOs []dtos.MangaDTO
	for _, Manga := range Manga {
		MangaDTOs = append(MangaDTOs, bs.mangaMapper.EntityToDTO(Manga))
	}

	return MangaDTOs, nil
}

func (bs MangaServiceImpl) CreateManga(MangaInsertDTO dtos.MangaInsertDTO) error {
	newManga := bs.mangaMapper.InsertDtoToEntity(MangaInsertDTO)

	if _, err := bs.mangaRepository.Create(context.TODO(), &newManga); err != nil {
		return err
	}

	return nil
}

func (bs MangaServiceImpl) UpdateManga(MangaId primitive.ObjectID, MangaInsertDTO dtos.MangaInsertDTO) error {
	mangaUpdated := bs.mangaMapper.InsertDtoToEntity(MangaInsertDTO)

	if _, err := bs.mangaRepository.Update(context.TODO(), MangaId, &mangaUpdated); err != nil {
		return err
	}

	return nil
}

func (bs MangaServiceImpl) DeleteManga(MangaId primitive.ObjectID) error {
	if _, err := bs.mangaRepository.DeleteById(context.TODO(), MangaId); err != nil {
		return err
	}
	return nil
}
