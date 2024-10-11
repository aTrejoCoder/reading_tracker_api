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

type BookService interface {
	GetBookId(bookId primitive.ObjectID) (*dtos.BookDTO, error)
	GetBookByISBN(ISBN string) (*dtos.BookDTO, error)
	GetBooksByNamePattern(bookName string, page int64, limit int64) ([]dtos.BookDTO, error)
	GetBooksByAuthor(author string, page int64, limit int64) ([]dtos.BookDTO, error)
	GetBooksByGenre(genre string, page int64, limit int64) ([]dtos.BookDTO, error)
	GetAllBooksSortedPaginated(page int64, limit int64) ([]dtos.BookDTO, error)

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

func (bs bookServiceImpl) GetBooksByNamePattern(bookName string, page int64, limit int64) ([]dtos.BookDTO, error) {
	filter := bson.M{"name": bson.M{"$regex": "^" + bookName, "$options": "i"}}

	books, err := bs.commonRepository.GetManyByFilterPaginated(context.TODO(), filter, page, limit)
	if err != nil {
		return nil, err
	}

	var bookDTOs []dtos.BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, bs.bookMapper.EntityToDTO(book))
	}
	return bookDTOs, nil
}

func (bs bookServiceImpl) GetBooksByAuthor(author string, page int64, limit int64) ([]dtos.BookDTO, error) {
	filter := bson.M{"author": bson.M{"$regex": author, "$options": "i"}}

	books, err := bs.commonRepository.GetManyByFilterPaginated(context.TODO(), filter, page, limit)
	if err != nil {
		return nil, err
	}

	var bookDTOs []dtos.BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, bs.bookMapper.EntityToDTO(book))
	}
	return bookDTOs, nil
}

func (bs bookServiceImpl) GetBooksByGenre(genre string, page int64, limit int64) ([]dtos.BookDTO, error) {
	filter := bson.M{"genres": bson.M{"$regex": genre, "$options": "i"}}

	books, err := bs.commonRepository.GetManyByFilterPaginated(context.TODO(), filter, page, limit)
	if err != nil {
		return nil, err
	}

	var bookDTOs []dtos.BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, bs.bookMapper.EntityToDTO(book))
	}
	return bookDTOs, nil
}

func (bs bookServiceImpl) GetBookByISBN(ISBN string) (*dtos.BookDTO, error) {
	filter := bson.M{"ISBN": ISBN}
	book, err := bs.commonRepository.GetByFilter(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	bookDTO := bs.bookMapper.EntityToDTO(*book)
	return &bookDTO, nil
}

func (bs bookServiceImpl) GetAllBooksSortedPaginated(page int64, limit int64) ([]dtos.BookDTO, error) {
	sortFields := bson.D{
		{Key: "name", Value: 1},
		{Key: "author", Value: 1},
	}

	books, err := bs.commonRepository.GetAllSortedPaginated(context.TODO(), sortFields, page, limit)
	if err != nil {
		return nil, err
	}

	var bookDTOs []dtos.BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, bs.bookMapper.EntityToDTO(book))
	}

	return bookDTOs, nil
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
	if _, err := bs.commonRepository.DeleteById(context.TODO(), bookId); err != nil {
		return err
	}
	return nil
}
