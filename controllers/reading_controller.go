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

type ReadingController struct {
	apiResponse    utils.ApiResponse
	validator      validator.Validate
	readingService services.ReadingService
}

func NewReadingControler(readingService services.ReadingService) *ReadingController {
	return &ReadingController{
		readingService: readingService,
		validator:      *validator.New(),
	}
}

func (c ReadingController) GetReadingById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		readingId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingDTOs, err := c.readingService.GetReadingById(readingId)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Found(ctx, readingDTOs, "Readings")
	}
}

func (c ReadingController) CreateReading() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var readingInsertDTO dtos.ReadingInsertDTO
		if isStructValid := utils.BindAndValidate(ctx, &readingInsertDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingService.CreateReading(readingInsertDTO); err != nil {
			c.apiResponse.ServerError(ctx, "Reading")
			return
		}

		c.apiResponse.Created(ctx, nil, "Reading")
	}
}

func (c ReadingController) UpdateReading() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		readingId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var readingInsertDTO dtos.ReadingInsertDTO
		if isStructValid := utils.BindAndValidate(ctx, &readingInsertDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingService.UpdateReading(readingId, readingInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Reading")
			return
		}

		c.apiResponse.Updated(ctx, nil, "Reading")
	}
}

func (c ReadingController) DeleteReading() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		err = c.readingService.DeleteReading(userId)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Deleted(ctx, "Reading")
	}
}
