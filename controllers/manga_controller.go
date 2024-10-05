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

type MangaController struct {
	mangaService services.MangaService
	apiResponse  utils.ApiResponse
	validator    *validator.Validate
}

func NewMangaController(mangaService services.MangaService) *MangaController {
	return &MangaController{
		mangaService: mangaService,
		validator:    validator.New(),
	}
}

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
