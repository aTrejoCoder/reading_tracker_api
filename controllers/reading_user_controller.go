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

// ReadingUserController handles user reading operations.
type ReadingUserController struct {
	apiResponse    utils.ApiResponse
	validator      validator.Validate
	readingService services.ReadingService
}

// NewReadingUserController creates a new ReadingUserController.
func NewReadingUserController(readingService services.ReadingService) *ReadingUserController {
	return &ReadingUserController{
		readingService: readingService,
		validator:      *validator.New(),
	}
}

// GetMyReadings retrieves all readings for the authenticated user.
// @Summary Get my readings
// @Description Retrieve a paginated list of readings for the authenticated user.
// @Tags Readings
// @Accept json
// @Produce json
// @Param sort query string false "Sort order (asc or desc)"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} dtos.ReadingDTO "Success"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/readings [get]
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

// GetMyReadingsByType retrieves readings for the authenticated user filtered by type.
// @Summary Get my readings by type
// @Description Retrieve readings for the authenticated user filtered by type.
// @Tags Readings
// @Accept json
// @Produce json
// @Param type query string false "Reading type (manga, book, custom_document)" default(book)
// @Param sort query string false "Sort order (asc or desc)"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} dtos.ReadingDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid reading type"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/readings/type [get]
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

// GetMyReadingsByStatus retrieves readings for the authenticated user filtered by status.
// @Summary Get my readings by status
// @Description Retrieve readings for the authenticated user filtered by status.
// @Tags Readings
// @Accept json
// @Produce json
// @Param status query string false "Reading status (ongoing, paused, completed)" default(ongoing)
// @Param sort query string false "Sort order (asc or desc)"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {array} dtos.ReadingDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid reading status"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/readings/status [get]
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

// StartReading initiates a new reading for the authenticated user.
// @Summary Start a reading
// @Description Start a new reading for the authenticated user.
// @Tags Readings
// @Accept json
// @Produce json
// @Param reading body dtos.ReadingInsertDTO true "Reading data"
// @Success 201 {object} utils.ApiResponse "Created"
// @Failure 400 {object} utils.ApiResponse "Validation error"
// @Failure 404 {object} utils.ApiResponse "Document not found"
// @Failure 409 {object} utils.ApiResponse "Duplicated reading"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/readings [post]
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
				c.apiResponse.Error(ctx, "The requested document already has reading", http.StatusBadRequest)
				return
			} else {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
		}

		c.apiResponse.Created(ctx, nil, "Reading")
	}
}

// UpdateMyReading updates an existing reading for the authenticated user.
// @Summary Update my reading
// @Description Update an existing reading for the authenticated user.
// @Tags Readings
// @Accept json
// @Produce json
// @Param readingId path string true "Reading ID"
// @Param reading body dtos.ReadingInsertDTO true "Updated reading data"
// @Success 200 {object} utils.ApiResponse "Updated"
// @Failure 400 {object} utils.ApiResponse "Validation error"
// @Failure 404 {object} utils.ApiResponse "Reading not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/readings/{readingId} [put]
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

// DeleteMyReading deletes a reading for the authenticated user.
// @Summary Delete my reading
// @Description Delete a reading for the authenticated user.
// @Tags Readings
// @Accept json
// @Produce json
// @Param readingId path string true "Reading ID"
// @Success 204 {object} utils.ApiResponse "Deleted"
// @Failure 404 {object} utils.ApiResponse "Reading not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/readings/{readingId} [delete]
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
