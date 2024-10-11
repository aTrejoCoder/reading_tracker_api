package controllers

import (
	"errors"
	"net/http"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecordUserController struct {
	recordService services.ReadingRecordService
	validator     *validator.Validate
	apiResponse   utils.ApiResponse
}

func NewRecordUserController(recordService services.ReadingRecordService) *RecordUserController {
	return &RecordUserController{
		recordService: recordService,
		validator:     validator.New(),
	}
}

// GetRecordsFromMyReading godoc
// @Summary Get records for a specific reading
// @Description Retrieve all records from a user's reading by reading ID
// @Tags ReadingRecord
// @Accept json
// @Produce json
// @Param readingId path string true "Reading ID"
// @Success 200 {object} utils.ApiResponse "Found reading records"
// @Failure 404 {object} utils.ApiResponse "No records found"
// @Failure 400 {object} utils.ApiResponse "Invalid request parameters"
// @Router /records/{readingId} [get]
func (c RecordUserController) GetRecordsFromMyReading() gin.HandlerFunc {
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

		recordsDTOs, err := c.recordService.GetRecordFromUser(readingId, userId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Reading Records")
			return
		}

		c.apiResponse.Found(ctx, recordsDTOs, "Reading Records")
	}
}

// AddRecord godoc
// @Summary Add a new record for a reading
// @Description Create a new record associated with a specific reading
// @Tags ReadingRecord
// @Accept json
// @Produce json
// @Param body body dtos.ReadingRecordInsertDTO true "Reading Record details"
// @Success 201 {object} utils.ApiResponse "Reading Record created"
// @Failure 400 {object} utils.ApiResponse "Invalid request body"
// @Failure 404 {object} utils.ApiResponse "Reading not found"
// @Router /records [post]
func (c RecordUserController) AddRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		var recordInserDTO dtos.ReadingRecordInsertDTO

		if isStructValid := utils.BindAndValidate(ctx, &recordInserDTO, c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.recordService.CreateRecord(recordInserDTO, userId); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
				return
			}

			c.apiResponse.NotFound(ctx, "Reading")
			return
		}

		c.apiResponse.Created(ctx, nil, "Reading Record")
	}
}

// UpdateRecord godoc
// @Summary Update an existing record for a reading
// @Description Update the details of a specific record for a reading
// @Tags ReadingRecord
// @Accept json
// @Produce json
// @Param readingId path string true "Reading ID"
// @Param body body dtos.ReadingRecordInsertDTO true "Updated Reading Record details"
// @Success 200 {object} utils.ApiResponse "Reading Record updated"
// @Failure 400 {object} utils.ApiResponse "Invalid request body"
// @Failure 404 {object} utils.ApiResponse "Reading not found"
// @Router /records/{readingId} [put]
func (c RecordUserController) UpdateRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		var recordInserDTO dtos.ReadingRecordInsertDTO

		if isStructValid := utils.BindAndValidate(ctx, &recordInserDTO, c.validator, c.apiResponse); !isStructValid {
			return
		}

		readingId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.recordService.UpdateUserRecord(readingId, userId, recordInserDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
				return
			}

			c.apiResponse.NotFound(ctx, "Reading")
			return
		}

		c.apiResponse.Created(ctx, nil, "Reading Record")
	}
}

// RemoveMyRecord godoc
// @Summary Remove a record from my readings
// @Description Delete a specific reading record from a user's readings
// @Tags ReadingRecord
// @Accept json
// @Produce json
// @Param readingId query string true "Reading ID"
// @Param recordId path string true "Record ID"
// @Success 200 {object} utils.ApiResponse "Reading Record deleted"
// @Failure 400 {object} utils.ApiResponse "Invalid request parameters"
// @Failure 404 {object} utils.ApiResponse "Reading not found"
// @Router /records/{recordId} [delete]
func (c RecordUserController) RemoveMyRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		readingIdStr := ctx.Query("readingId")
		if readingIdStr == "" {
			c.apiResponse.Error(ctx, "readingId is empty", http.StatusBadRequest)
			return
		}

		readingId, err := primitive.ObjectIDFromHex(readingIdStr)
		if err != nil {
			c.apiResponse.Error(ctx, "readingId must be an objectId", http.StatusBadRequest)
			return
		}

		recordId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.recordService.DeleteRecord(readingId, userId, recordId); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.Error(ctx, err.Error(), http.StatusInternalServerError)
				return
			}

			c.apiResponse.NotFound(ctx, "Reading")
			return
		}

		c.apiResponse.Deleted(ctx, "Record")
	}
}
