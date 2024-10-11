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

type ReadingListUserController struct {
	readingListService services.ReadingListService
	apiResponse        utils.ApiResponse
	validator          validator.Validate
}

func NewReadingListUserController(readingListService services.ReadingListService) *ReadingListUserController {
	return &ReadingListUserController{
		readingListService: readingListService,
		validator:          *validator.New(),
	}
}

// GetMyReadingLists godoc
// @Summary Get reading lists for the authenticated user
// @Description Retrieve all reading lists associated with the authenticated user
// @Tags ReadingList
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse "Successful retrieval"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Router /my-reading-lists [get]
func (c ReadingListUserController) GetMyReadingLists() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
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

// GetMyReadingListById godoc
// @Summary Get a specific reading list by ID for the authenticated user
// @Description Retrieve a reading list associated with the authenticated user by list ID
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param listId query string true "Reading List ID"
// @Success 200 {object} utils.ApiResponse "Successful retrieval"
// @Failure 400 {object} utils.ApiResponse "Invalid list ID"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /my-reading-lists [get]
func (c ReadingListUserController) GetMyReadingListById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		listId, err := utils.GetObjectIdFromUrlQuery(ctx, "listId")
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		readingListDTOs, err := c.readingListService.GetReadingListById(listId, userId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}
			c.apiResponse.NotFound(ctx, "Reading List")
			return
		}

		c.apiResponse.Found(ctx, readingListDTOs, "Reading List")
	}
}

// AddReadingToList godoc
// @Summary Add readings to a user's reading list
// @Description Add readings to the authenticated user's specified reading list
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param body body dtos.UpdateReadingsToListDTO true "Readings to be added"
// @Success 200 {object} utils.ApiResponse "Successful addition"
// @Failure 400 {object} utils.ApiResponse "Invalid request body"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /my-reading-lists [post]
func (c ReadingListUserController) AddReadingToList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		var insertReadingsToListDTO dtos.UpdateReadingsToListDTO
		if isStructValid := utils.BindAndValidate(ctx, &insertReadingsToListDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingListService.AddReadingToList(userId, insertReadingsToListDTO); err != nil {
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
// @Description Remove readings from the authenticated user's specified reading list
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param body body dtos.UpdateReadingsToListDTO true "Readings to be removed"
// @Success 200 {object} utils.ApiResponse "Successful removal"
// @Failure 400 {object} utils.ApiResponse "Invalid request body"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /my-reading-lists [delete]
func (c ReadingListUserController) RemoveReadingToList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		var insertReadingsToListDTO dtos.UpdateReadingsToListDTO
		if isStructValid := utils.BindAndValidate(ctx, &insertReadingsToListDTO, &c.validator, c.apiResponse); !isStructValid {
			return
		}

		if err := c.readingListService.RemoveReadingToList(userId, insertReadingsToListDTO); err != nil {
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
// @Summary Create a new reading list for the authenticated user
// @Description Create a new reading list associated with the authenticated user
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param body body dtos.ReadingListInsertDTO true "Reading List details"
// @Success 200 {object} utils.ApiResponse "Successful creation"
// @Failure 400 {object} utils.ApiResponse "Invalid request body"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Router /my-reading-lists [post]
func (c ReadingListUserController) CreateReadingList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
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

		c.apiResponse.Created(ctx, nil, "ReadingList")
	}
}

// UpdateMyReadingList godoc
// @Summary Update an existing reading list for the authenticated user
// @Description Update the details of a specific reading list associated with the authenticated user
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param readingId path string true "Reading List ID"
// @Param body body dtos.ReadingListInsertDTO true "Updated Reading List details"
// @Success 200 {object} utils.ApiResponse "Successful update"
// @Failure 400 {object} utils.ApiResponse "Invalid request body"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /my-reading-lists/{readingId} [put]
func (c ReadingListUserController) UpdateMyReadingList() gin.HandlerFunc {
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

// DeleteMyReadingList godoc
// @Summary Delete a reading list for the authenticated user
// @Description Delete a specific reading list associated with the authenticated user
// @Tags ReadingList
// @Accept json
// @Produce json
// @Param readingId path string true "Reading List ID"
// @Success 200 {object} utils.ApiResponse "Successful deletion"
// @Failure 404 {object} utils.ApiResponse "Reading List not found"
// @Router /my-reading-lists/{readingId} [delete]
func (c ReadingListUserController) DeleteMyReadingList() gin.HandlerFunc {
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
