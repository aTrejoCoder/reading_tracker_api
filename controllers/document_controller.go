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
	DocumentService services.DocumentService
	apiResponse     utils.ApiResponse
	validator       *validator.Validate
}

func NewDocumentController(DocumentService services.DocumentService) *DocumentController {
	return &DocumentController{
		DocumentService: DocumentService,
		validator:       validator.New(),
	}
}

func (c DocumentController) GetDocumentById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		DocumentId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		DocumentDTO, err := c.DocumentService.GetDocumentById(DocumentId)
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
		var DocumentInsertDTO dtos.DocumentInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &DocumentInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.DocumentService.CreateDocument(DocumentInsertDTO); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "Document")
	}
}

func (c DocumentController) UpdateDocument() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		documentId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		var DocumentInsertDTO dtos.DocumentInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &DocumentInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.DocumentService.UpdateDocument(documentId, DocumentInsertDTO); err != nil {
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
		DocumentId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		if err := c.DocumentService.DeleteDocument(DocumentId); err != nil {
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
