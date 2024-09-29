package controllers

import (
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MangaController struct {
	MangaService services.MangaService
	apiResponse  utils.ApiResponse
	validator    *validator.Validate
}

func NewMangaController(MangaService services.MangaService) *MangaController {
	return &MangaController{
		MangaService: MangaService,
		validator:    validator.New(),
	}
}

func (c MangaController) GetMangaById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mangaId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		mangaDTO, err := c.MangaService.GetMangaId(mangaId)
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

func (c MangaController) CreateManga() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var mangaInsertDTO dtos.MangaInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &mangaInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.MangaService.CreateManga(mangaInsertDTO); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "Manga")
	}
}

func (c MangaController) UpdateManga() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mangaId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		var MangaInsertDTO dtos.MangaInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &MangaInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.MangaService.UpdateManga(mangaId, MangaInsertDTO); err != nil {
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
		mangaId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		if err := c.MangaService.DeleteManga(mangaId); err != nil {
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
