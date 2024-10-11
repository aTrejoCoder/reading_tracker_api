package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
)

type MangaMapper struct {
}

func (mm MangaMapper) InsertDtoToEntity(mangaInsertDTO dtos.MangaInsertDTO) models.Manga {
	now := time.Now().UTC()

	return models.Manga{
		Title:           mangaInsertDTO.Title,
		Author:          mangaInsertDTO.Author,
		CoverImageURL:   mangaInsertDTO.CoverImageURL,
		Volume:          mangaInsertDTO.Volume,
		Chapters:        mangaInsertDTO.Chapters,
		Demography:      mangaInsertDTO.Demogragphy,
		Genres:          mangaInsertDTO.Genres,
		PublicationDate: mangaInsertDTO.PublicationDate,
		Publisher:       mangaInsertDTO.Publisher,
		Description:     mangaInsertDTO.Description,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (mm MangaMapper) EntityToDTO(manga models.Manga) dtos.MangaDTO {
	return dtos.MangaDTO{
		Id:              manga.Id,
		Title:           manga.Title,
		Author:          manga.Author,
		CoverImageURL:   manga.CoverImageURL,
		Volume:          manga.Volume,
		Chapters:        manga.Chapters,
		Demogragphy:     manga.Demography,
		Genres:          manga.Genres,
		PublicationDate: manga.PublicationDate,
		Publisher:       manga.Publisher,
		Description:     manga.Description,
	}
}
