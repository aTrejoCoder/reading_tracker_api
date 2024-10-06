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

type ReadingListController struct {
	readingListService services.ReadingListService
	apiResponse        utils.ApiResponse
	validator          validator.Validate
}

func NewReadingListController(readingListService services.ReadingListService) *ReadingListController {
	return &ReadingListController{
		readingListService: readingListService,
		validator:          *validator.New(),
	}
}

func (c ReadingListController) GetReadingListByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlQuery(ctx, "userId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingListDTOs, err := c.readingListService.GetReadingListsByUserId(userId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "User")
			return
		}

		c.apiResponse.Found(ctx, readingListDTOs, "Reading Lists")
	}
}

func (c ReadingListController) AddReadingToList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var insertReadingsToListDTO dtos.UpdateReadingsToListDTO
		if isStructValid := utils.BindAndValidate(ctx, &insertReadingsToListDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err = c.readingListService.AddReadingToList(userId, insertReadingsToListDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "Reading List")
			return
		}

		c.apiResponse.OK(ctx, nil, "Readings Succesfully Added")
	}
}

func (c ReadingListController) RemoveReadingToList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var insertReadingsToListDTO dtos.UpdateReadingsToListDTO
		if isStructValid := utils.BindAndValidate(ctx, &insertReadingsToListDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err = c.readingListService.RemoveReadingToList(userId, insertReadingsToListDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "Reading List")
			return
		}

		c.apiResponse.OK(ctx, nil, "Readings Succesfully Removed")
	}
}

func (c ReadingListController) CreateReadingList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlQuery(ctx, "userId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var readingListInsertDTO dtos.ReadingListInsertDTO
		if isStructValid := utils.BindAndValidate(ctx, &readingListInsertDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingListService.CreateReadingList(userId, readingListInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "Reading List")
			return
		}

		c.apiResponse.Updated(ctx, nil, "ReadingList")
	}
}

func (c ReadingListController) UpdateReadingList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlQuery(ctx, "userId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var readingListInsertDTO dtos.ReadingListInsertDTO
		if isStructValid := utils.BindAndValidate(ctx, &readingListInsertDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingListService.UpdateReadingList(readingId, userId, readingListInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "Reading List")
			return
		}

		c.apiResponse.Updated(ctx, nil, "ReadingList")
	}
}

func (c ReadingListController) DeleteReadingList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlQuery(ctx, "userId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		err = c.readingListService.DeleteReadingList(userId, readingId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "Reading List")
			return
		}

		c.apiResponse.Deleted(ctx, "Reading Lists")
	}
}
