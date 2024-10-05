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

type ReadingUserController struct {
	apiResponse    utils.ApiResponse
	validator      validator.Validate
	readingService services.ReadingService
}

func NewReadingUserController(readingService services.ReadingService) *ReadingUserController {
	return &ReadingUserController{
		readingService: readingService,
		validator:      *validator.New(),
	}
}

func (c ReadingUserController) GetMyReadings() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingDTOs, err := c.readingService.GetReadingsByUserId(userId)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Found(ctx, readingDTOs, "Readings")
	}
}

func (c ReadingUserController) GetMyMangaReadings() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingDTOs, err := c.readingService.GetReadingsByUserId(userId)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Found(ctx, readingDTOs, "Readings")
	}
}

func (c ReadingUserController) GetMyBookReadings() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingDTOs, err := c.readingService.GetReadingsByUserId(userId)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Found(ctx, readingDTOs, "Readings")
	}
}

func (c ReadingUserController) StartReading() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		var readingInsertDTO dtos.ReadingInsertDTO
		if isStructValid := utils.BindAndValidate(ctx, &readingInsertDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingService.CreateReading(readingInsertDTO, userId); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "Reading")
	}
}

func (c ReadingUserController) UpdateMyReading() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var readingInsertDTO dtos.ReadingInsertDTO
		if isStructValid := utils.BindAndValidate(ctx, &readingInsertDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingService.UpdateReading(readingId, userId, readingInsertDTO); err != nil {
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

func (c ReadingUserController) DeleteMyReading() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		err := c.readingService.DeleteReading(readingId, userId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Reading")
			return
		}

		c.apiResponse.Deleted(ctx, "Reading")
	}
}
