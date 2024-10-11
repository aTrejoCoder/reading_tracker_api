package mappers

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookMapper struct {
}

func (bm BookMapper) InsertDtoToEntity(insertDTO dtos.BookInsertDTO) models.Book {
	now := time.Now().UTC()
	return models.Book{
		Author:          insertDTO.Author,
		ISBN:            insertDTO.ISBN,
		Name:            insertDTO.Name,
		CoverImageURL:   insertDTO.CoverImageURL,
		Edition:         insertDTO.Edition,
		Pages:           insertDTO.Pages,
		Language:        insertDTO.Language,
		PublicationDate: insertDTO.PublicationDate,
		Publisher:       insertDTO.Publisher,
		Description:     insertDTO.Description,
		Genres:          insertDTO.Genres,
		ReadingList:     []primitive.ObjectID{},
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

func (bm BookMapper) EntityToDTO(book models.Book) dtos.BookDTO {
	return dtos.BookDTO{
		Id:              book.Id,
		Author:          book.Author,
		ISBN:            book.ISBN,
		Name:            book.Name,
		CoverImageURL:   book.CoverImageURL,
		Edition:         book.Edition,
		Pages:           book.Pages,
		Language:        book.Language,
		PublicationDate: book.PublicationDate,
		Publisher:       book.Publisher,
		Description:     book.Description,
		Genres:          book.Genres,
	}
}
