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
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		sortParam := ctx.DefaultQuery("sort", "asc")
		isAsc := sortParam != "desc"

		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingDTOs, err := c.readingService.GetReadingsByUserId(userId, page, limit, isAsc)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Found(ctx, readingDTOs, "Readings")
	}
}

func (c ReadingUserController) GetMyReadingsByType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		sortParam := ctx.DefaultQuery("sort", "asc")
		isAsc := sortParam != "desc"

		// Valid reading types (manga or book)
		readingTypeParam := ctx.DefaultQuery("type", "book")
		validTypes := map[string]bool{
			"manga":           true,
			"book":            true,
			"custom_document": true,
		}

		if !validTypes[readingTypeParam] {
			c.apiResponse.Error(ctx, "invalid reading type", http.StatusBadRequest)
			return
		}

		// Sorted by Record Updates
		sortOrder := "last_record_update"

		readingDTOs, err := c.readingService.GetReadingsByUserAndType(userId, sortOrder, readingTypeParam, page, limit, isAsc)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		c.apiResponse.Found(ctx, readingDTOs, "Readings")
	}
}

func (c ReadingUserController) GetMyReadingsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, limit := utils.GetPaginationValuesFromRequest(ctx)

		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		sortParam := ctx.DefaultQuery("sort", "asc")
		isAsc := sortParam != "desc"

		statusParam := ctx.DefaultQuery("status", "ongoing")
		validStatuses := map[string]bool{
			"ongoing":   true,
			"paused":    true,
			"completed": true,
		}

		if !validStatuses[statusParam] {
			c.apiResponse.Error(ctx, "invalid reading status", http.StatusBadRequest)
			return
		}

		readingDTOs, err := c.readingService.GetReadingsByUserAndStatus(userId, statusParam, page, limit, isAsc)
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
			if errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.NotFound(ctx, "document")
				return
			} else if errors.Is(err, utils.ErrDuplicated) {
				c.apiResponse.Error(ctx, "The requested document aleady has reading", http.StatusBadRequest)
				return
			} else {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
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
