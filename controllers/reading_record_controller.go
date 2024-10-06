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

type RecordController struct {
	recordService services.ReadingRecordService
	validator     *validator.Validate
	apiResponse   utils.ApiResponse
}

func NewRecordController(recordService services.ReadingRecordService) *RecordController {
	return &RecordController{
		recordService: recordService,
		validator:     validator.New(),
	}
}

func (c RecordController) CreateRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlQuery(ctx, "userId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
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

func (c RecordController) UpdateRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recordId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var recordInserDTO dtos.ReadingRecordInsertDTO

		if isStructValid := utils.BindAndValidate(ctx, &recordInserDTO, c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.recordService.UpdateRecord(recordId, recordInserDTO); err != nil {
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

func (c RecordController) DeleteRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		recordId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingIdStr := ctx.Query("readingId")
		if readingIdStr == "" {
			c.apiResponse.Error(ctx, "readingId is empty", http.StatusBadRequest)
			return
		}

		userId, err := utils.GetObjectIdFromUrlQuery(ctx, "userId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingId, err := primitive.ObjectIDFromHex(readingIdStr)
		if err != nil {
			c.apiResponse.Error(ctx, "readingId must be an objectId", http.StatusBadRequest)
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
