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

// MangaController handles manga-related operations.
type MangaController struct {
	mangaService services.MangaService
	apiResponse  utils.ApiResponse
	validator    *validator.Validate
}

// NewMangaController creates a new MangaController.
func NewMangaController(mangaService services.MangaService) *MangaController {
	return &MangaController{
		mangaService: mangaService,
		validator:    validator.New(),
	}
}

// GetMangaById retrieves a manga by ID.
// @Summary Get Manga by ID
// @Description Retrieve a manga by its ID.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param mangaId path string true "Manga ID"
// @Success 200 {object} dtos.MangaDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid Manga ID"
// @Failure 404 {object} utils.ApiResponse "Manga not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/{mangaId} [get]
func (c MangaController) GetMangaById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mangaId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		mangaDTO, err := c.mangaService.GetMangaById(mangaId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Manga")
			return
		}

		c.apiResponse.Found(ctx, mangaDTO, "Manga")
	}
}

// GetMangaByAuthor retrieves mangas by a specific author.
// @Summary Get Mangas by Author
// @Description Retrieve all mangas written by a specific author.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param author path string true "Author Name"
// @Success 200 {array} dtos.MangaDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid Author Name"
// @Failure 404 {object} utils.ApiResponse "No mangas found for this author"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/author/{author} [get]
func (c MangaController) GetMangaByAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		author := ctx.Param("author")
		if author == "" {
			c.apiResponse.Error(ctx, "Author not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.mangaService.GetMangaByAuthor(author)
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

// GetMangaByDemography retrieves mangas by demography.
// @Summary Get Mangas by Demography
// @Description Retrieve all mangas under a specific demography.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param demography path string true "Demography Type"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of results per page" default(10)
// @Success 200 {array} dtos.MangaDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid Demography"
// @Failure 404 {object} utils.ApiResponse "No mangas found for this demography"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/demography/{demography} [get]
func (c MangaController) GetMangaByDemography() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		demography := ctx.Param("demography")
		if demography == "" {
			c.apiResponse.Error(ctx, "Demography not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.mangaService.GetMangaByDemography(demography, page, limit)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Demography")
			return
		}

		c.apiResponse.Found(ctx, booksDTO, "Demography")
	}
}

// GetAllMangasSortedPaginated retrieves all mangas with sorting and pagination.
// @Summary Get All Mangas Sorted and Paginated
// @Description Retrieve all mangas with sorting and pagination options.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of results per page" default(10)
// @Success 200 {array} dtos.MangaDTO "Success"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas [get]
func (c MangaController) GetAllMangasSortedPaginated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		bookDTOs, err := c.mangaService.GetAllMangaSortedPaginated(page, limit)
		if err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Found(ctx, bookDTOs, "Books")
	}
}

// GetMangaByGenre retrieves mangas by genre.
// @Summary Get Mangas by Genre
// @Description Retrieve all mangas under a specific genre.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param genre path string true "Genre"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of results per page" default(10)
// @Success 200 {array} dtos.MangaDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid Genre"
// @Failure 404 {object} utils.ApiResponse "No mangas found for this genre"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/genre/{genre} [get]
func (c MangaController) GetMangaByGenre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		genre := ctx.Param("genre")
		if genre == "" {
			c.apiResponse.Error(ctx, "Genre not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.mangaService.GetMangaByGenre(genre, page, limit)
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

// GetMangaByMatchingName retrieves mangas by name pattern.
// @Summary Get Mangas by Name Pattern
// @Description Retrieve all mangas that match a specific name pattern.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param name path string true "Manga Name"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of results per page" default(10)
// @Success 200 {array} dtos.MangaDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid Name"
// @Failure 404 {object} utils.ApiResponse "No mangas found for this name"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/name/{name} [get]
func (c MangaController) GetMangaByMatchingName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		bookName := ctx.Param("name")
		if bookName == "" {
			c.apiResponse.Error(ctx, "Book name not provided", http.StatusBadRequest)
			return
		}

		booksDTO, err := c.mangaService.GetMangaByNamePattern(bookName, page, limit)
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

// CreateManga creates a new manga.
// @Summary Create Manga
// @Description Create a new manga in the system.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param mangaInsertDTO body dtos.MangaInsertDTO true "Manga Insert DTO"
// @Success 201 {object} utils.ApiResponse "Manga Created"
// @Failure 400 {object} utils.ApiResponse "Invalid Data"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas [post]
func (c MangaController) CreateManga() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mangaInsertDTO dtos.MangaInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &mangaInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.mangaService.CreateManga(mangaInsertDTO); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "Manga")
	}
}

// UpdateManga updates an existing manga.
// @Summary Update Manga
// @Description Update an existing manga by its ID.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param mangaId path string true "Manga ID"
// @Param mangaInsertDTO body dtos.MangaInsertDTO true "Manga Insert DTO"
// @Success 200 {object} utils.ApiResponse "Manga Updated"
// @Failure 400 {object} utils.ApiResponse "Invalid Manga ID"
// @Failure 404 {object} utils.ApiResponse "Manga not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/{mangaId} [put]
func (c MangaController) UpdateManga() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mangaId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		var MangaInsertDTO dtos.MangaInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &MangaInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.mangaService.UpdateManga(mangaId, MangaInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Manga")
			return
		}

		c.apiResponse.Updated(ctx, nil, "Manga")
	}
}

// DeleteManga deletes a manga by ID.
// @Summary Delete Manga
// @Description Delete a manga by its ID.
// @Tags Mangas
// @Accept json
// @Produce json
// @Param mangaId path string true "Manga ID"
// @Success 200 {object} utils.ApiResponse "Manga Deleted"
// @Failure 400 {object} utils.ApiResponse "Invalid Manga ID"
// @Failure 404 {object} utils.ApiResponse "Manga not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/mangas/{mangaId} [delete]
func (c MangaController) DeleteManga() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mangaId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		if err := c.mangaService.DeleteManga(mangaId); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Manga")
			return
		}

		c.apiResponse.Deleted(ctx, "Manga")
	}
}
