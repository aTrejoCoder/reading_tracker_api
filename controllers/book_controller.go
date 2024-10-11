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

// BookController handles book-related operations
type BookController struct {
	bookService services.BookService
	apiResponse utils.ApiResponse
	validator   *validator.Validate
}

// NewBookController creates a new instance of BookController
func NewBookController(bookService services.BookService) *BookController {
	return &BookController{
		bookService: bookService,
		validator:   validator.New(),
	}
}

// GetBookById retrieves a book by its ID
// @Summary Get a book by ID
// @Description Get book details by ID
// @Tags Books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} dtos.BookDTO
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /books/{id} [get]
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

// GetBookByISBN retrieves a book by its ISBN
// @Summary Get a book by ISBN
// @Description Get book details by ISBN
// @Tags Books
// @Produce json
// @Param isbn path string true "Book ISBN"
// @Success 200 {object} dtos.BookDTO
// @Failure 404 {object} utils.ApiResponse
// @Router /books/isbn/{isbn} [get]
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

// GetBooksByAuthor retrieves books by the author's name
// @Summary Get books by author
// @Description Get books written by a specific author
// @Tags Books
// @Produce json
// @Param author path string true "Author Name"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of books per page" default(10)
// @Success 200 {object} []dtos.BookDTO
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /books/author/{author} [get]
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

// GetAllBooksSortedPaginated retrieves all books with sorting and pagination
// @Summary Get all books with sorting and pagination
// @Description Get all books sorted and paginated
// @Tags Books
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of books per page" default(10)
// @Success 200 {object} []dtos.BookDTO
// @Failure 500 {object} utils.ApiResponse
// @Router /books [get]
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

// GetBooksByGenre retrieves books by genre
// @Summary Get books by genre
// @Description Get books of a specific genre
// @Tags Books
// @Produce json
// @Param genre path string true "Genre"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of books per page" default(10)
// @Success 200 {object} []dtos.BookDTO
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /books/genre/{genre} [get]
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

// GetBooksByMatchingName retrieves books matching a given name
// @Summary Get books by name pattern
// @Description Get books matching a specific name
// @Tags Books
// @Produce json
// @Param name path string true "Book Name"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of books per page" default(10)
// @Success 200 {object} []dtos.BookDTO
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /books/name/{name} [get]
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

// CreateBook creates a new book
// @Summary Create a new book
// @Description Create a new book entry
// @Tags Books
// @Accept json
// @Produce json
// @Param book body dtos.BookInsertDTO true "Book Information"
// @Success 201 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /books [post]
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

// UpdateBook updates an existing book
// @Summary Update an existing book
// @Description Update a book entry by ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body dtos.BookInsertDTO true "Book Information"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /books/{id} [put]
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

// DeleteBook deletes a book by its ID
// @Summary Delete a book
// @Description Delete a book entry by ID
// @Tags Books
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /books/{id} [delete]
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
