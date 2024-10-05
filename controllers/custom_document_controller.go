package controllers

import (
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DocumentController struct {
	CustomDocumentService services.CustomDocumentService
	apiResponse           utils.ApiResponse
	validator             *validator.Validate
}

func NewDocumentController(CustomDocumentService services.CustomDocumentService) *DocumentController {
	return &DocumentController{
		CustomDocumentService: CustomDocumentService,
		validator:             validator.New(),
	}
}

func (c DocumentController) GetDocumentById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		documentId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		DocumentDTO, err := c.CustomDocumentService.GetCustomDocumentById(userId, documentId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Document")
			return
		}

		c.apiResponse.Found(ctx, DocumentDTO, "Document")
	}
}

func (c DocumentController) GetMyCustomDocuments() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		DocumentDTO, err := c.CustomDocumentService.GetCustomDocumentsByUserId(userId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Document")
			return
		}

		c.apiResponse.Found(ctx, DocumentDTO, "Document")
	}
}

func (c DocumentController) CreateDocument() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		var documentInsertDTO dtos.CustomDocumentInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &documentInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.CustomDocumentService.CreateCustomDocument(documentInsertDTO, userId); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "Document")
	}
}

func (c DocumentController) UpdateDocument() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		documentId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		var documentInsertDTO dtos.CustomDocumentInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &documentInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.CustomDocumentService.UpdateCustomDocument(documentId, userId, documentInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Document")
			return
		}

		c.apiResponse.Updated(ctx, nil, "Document")
	}
}

func (c DocumentController) DeleteDocument() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, isUserIdRetrieved := utils.GetUserIdFromRequest(ctx, c.apiResponse)
		if !isUserIdRetrieved {
			return
		}

		documentId, err := utils.GetObjectIdFromUrlParam(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		if err := c.CustomDocumentService.DeleteCustomDocument(userId, documentId); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Document")
			return
		}

		c.apiResponse.Deleted(ctx, "Document")
	}
}
