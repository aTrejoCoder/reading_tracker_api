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

		c.apiResponse.OK(ctx, nil, "Readings Succesfully Added")
	}
}

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

		c.apiResponse.OK(ctx, nil, "Readings Succesfully Removed")
	}
}

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
