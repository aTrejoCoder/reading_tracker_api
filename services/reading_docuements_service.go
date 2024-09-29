package services

import (
	"context"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/mappers"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService interface {
	GetBookId(bookId primitive.ObjectID) (*dtos.BookDTO, error)
	CreateBook(bookInsertDTO dtos.BookInsertDTO) error
	UpdateBook(bookId primitive.ObjectID, bookInsertDTO dtos.BookInsertDTO) error
	DeleteBook(bookId primitive.ObjectID) error
}

type bookServiceImpl struct {
	bookMapper       mappers.BookMapper
	commonRepository repository.Repository[models.Book]
}

func NewBookService(commonRepository repository.Repository[models.Book]) BookService {
	return &bookServiceImpl{
		commonRepository: commonRepository,
	}

}

func (bs bookServiceImpl) GetBookId(bookId primitive.ObjectID) (*dtos.BookDTO, error) {
	book, err := bs.commonRepository.GetByID(context.TODO(), bookId)
	if err != nil {
		return nil, err
	}

	bookDTO := bs.bookMapper.EntityToDTO(*book)
	return &bookDTO, nil
}

func (bs bookServiceImpl) CreateBook(bookInsertDTO dtos.BookInsertDTO) error {
	newBook := bs.bookMapper.InsertDtoToEntity(bookInsertDTO)

	if _, err := bs.commonRepository.Create(context.TODO(), &newBook); err != nil {
		return err
	}

	return nil
}

func (bs bookServiceImpl) UpdateBook(bookId primitive.ObjectID, bookInsertDTO dtos.BookInsertDTO) error {
	newBook := bs.bookMapper.InsertDtoToEntity(bookInsertDTO)

	if _, err := bs.commonRepository.Update(context.TODO(), bookId, &newBook); err != nil {
		return err
	}

	return nil
}

func (bs bookServiceImpl) DeleteBook(bookId primitive.ObjectID) error {
	if _, err := bs.commonRepository.Delete(context.TODO(), bookId); err != nil {
		return err
	}
	return nil
}
