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

// GetReadingListByUserId godoc
// @Summary Get reading lists by user ID
// @Description Retrieve all reading lists associated with a specific user ID
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Success 200 {object} utils.ApiResponse "Successful retrieval"
// @Failure 400 {object} utils.ApiResponse "Invalid user ID"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Router /reading-lists [get]
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

// AddReadingToList godoc
// @Summary Add readings to a user's reading list
// @Description Add readings to the specified user's reading list
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param body body dtos.UpdateReadingsToListDTO true "Readings to be added"
// @Success 200 {object} utils.ApiResponse "Successful addition"
// @Failure 400 {object} utils.ApiResponse "Invalid user ID or request body"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /reading-lists/{userId} [post]
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

		c.apiResponse.OK(ctx, nil, "Readings Successfully Added")
	}
}

// RemoveReadingToList godoc
// @Summary Remove readings from a user's reading list
// @Description Remove readings from the specified user's reading list
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param body body dtos.UpdateReadingsToListDTO true "Readings to be removed"
// @Success 200 {object} utils.ApiResponse "Successful removal"
// @Failure 400 {object} utils.ApiResponse "Invalid user ID or request body"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /reading-lists/{userId} [delete]
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

		c.apiResponse.OK(ctx, nil, "Readings Successfully Removed")
	}
}

// CreateReadingList godoc
// @Summary Create a new reading list
// @Description Create a new reading list for a specific user
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Param body body dtos.ReadingListInsertDTO true "Reading List details"
// @Success 200 {object} utils.ApiResponse "Successful creation"
// @Failure 400 {object} utils.ApiResponse "Invalid user ID or request body"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Router /reading-lists [post]
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

		c.apiResponse.Updated(ctx, nil, "Reading List")
	}
}

// UpdateReadingList godoc
// @Summary Update an existing reading list
// @Description Update the details of a specific reading list
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Param readingId path string true "Reading List ID"
// @Param body body dtos.ReadingListInsertDTO true "Updated Reading List details"
// @Success 200 {object} utils.ApiResponse "Successful update"
// @Failure 400 {object} utils.ApiResponse "Invalid user ID, reading list ID, or request body"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /reading-lists/{readingId} [put]
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

		c.apiResponse.Updated(ctx, nil, "Reading List")
	}
}

// DeleteReadingList godoc
// @Summary Delete a reading list
// @Description Delete a specific reading list associated with a user
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param userId query string true "User ID"
// @Param readingId path string true "Reading List ID"
// @Success 200 {object} utils.ApiResponse "Successful deletion"
// @Failure 400 {object} utils.ApiResponse "Invalid user ID or reading list ID"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /reading-lists/{readingId} [delete]
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
