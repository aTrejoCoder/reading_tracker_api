package controllers

import (
	"errors"
	"net/http"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookController struct {
	bookService services.BookService
	apiResponse utils.ApiResponse
	validator   *validator.Validate
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{
		bookService: bookService,
		validator:   validator.New(),
	}
}

func (c BookController) GetBookById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		bookDTO, err := c.bookService.GetBookId(bookId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Book")
			return
		}

		c.apiResponse.Found(ctx, bookDTO, "Book")
	}
}

func (c BookController) GetBookByISBN() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookISBN := ctx.Param("isbn")
		bookDTO, err := c.bookService.GetBookByISBN(bookISBN)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Book")
			return
		}

		c.apiResponse.Found(ctx, bookDTO, "Book")
	}
}

func (c BookController) GetBooksByAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		author := ctx.Param("author")
		if author == "" {
			c.apiResponse.Error(ctx, "Author not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.bookService.GetBooksByAuthor(author, page, limit)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Author")
			return
		}

		c.apiResponse.Found(ctx, booksDTO, "Author Books")
	}
}

func (c BookController) GetAllBooksSortedPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		bookDTOs, err := c.bookService.GetAllBooksSortedPaginated(page, limit)
		if err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Found(ctx, bookDTOs, "Books")
	}
}

func (c BookController) GetBooksByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		genre := ctx.Param("genre")
		if genre == "" {
			c.apiResponse.Error(ctx, "Genre not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.bookService.GetBooksByGenre(genre, page, limit)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Genre")
			return
		}

		c.apiResponse.Found(ctx, booksDTO, "Books By Genre")
	}
}

func (c BookController) GetBooksByMatchingName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		bookName := ctx.Param("name")
		if bookName == "" {
			c.apiResponse.Error(ctx, "Book name not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.bookService.GetBooksByNamePattern(bookName, page, limit)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Book")
			return
		}

		c.apiResponse.Found(ctx, booksDTO, "Book")
	}
}

func (c BookController) CreateBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookInsertDTO dtos.BookInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &bookInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.bookService.CreateBook(bookInsertDTO); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "book")
	}
}

func (c BookController) UpdateBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		var bookInsertDTO dtos.BookInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &bookInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.bookService.UpdateBook(bookId, bookInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Book")
			return
		}

		c.apiResponse.Updated(ctx, nil, "book")
	}
}

func (c BookController) DeleteBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		if err := c.bookService.DeleteBook(bookId); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Book")
			return
		}

		c.apiResponse.Deleted(ctx, "book")
	}
}
