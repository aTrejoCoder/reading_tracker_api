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
