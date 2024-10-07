package controllers

import (
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// DocumentController handles requests related to custom documents.
type DocumentController struct {
	CustomDocumentService services.CustomDocumentService
	apiResponse           utils.ApiResponse
	validator             *validator.Validate
}

// NewDocumentController creates a new DocumentController.
func NewDocumentController(CustomDocumentService services.CustomDocumentService) *DocumentController {
	return &DocumentController{
		CustomDocumentService: CustomDocumentService,
		validator:             validator.New(),
	}
}

// GetDocumentById retrieves a document by its ID.
// @Summary Get a document by ID
// @Description Retrieve a custom document by its ID for the authenticated user
// @Tags documents
// @Produce json
// @Param id path string true "Document ID"
// @Success 200 {object} dtos.CustomDocumentDTO
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /documents/{id} [get]
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

// GetMyCustomDocuments retrieves all custom documents for the user.
// @Summary Get all custom documents for the authenticated user
// @Description Retrieve all custom documents for the user
// @Tags documents
// @Produce json
// @Success 200 {array} dtos.CustomDocumentDTO
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /documents [get]
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

// CreateDocument creates a new custom document.
// @Summary Create a new custom document
// @Description Create a custom document for the authenticated user
// @Tags documents
// @Accept json
// @Produce json
// @Param document body dtos.CustomDocumentInsertDTO true "Custom Document Data"
// @Success 201 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /documents [post]
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

// UpdateDocument updates an existing custom document.
// @Summary Update an existing custom document
// @Description Update a custom document by ID for the authenticated user
// @Tags documents
// @Accept json
// @Produce json
// @Param id path string true "Document ID"
// @Param document body dtos.CustomDocumentInsertDTO true "Updated Custom Document Data"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /documents/{id} [put]
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

// DeleteDocument deletes a custom document.
// @Summary Delete a custom document
// @Description Delete a custom document by ID for the authenticated user
// @Tags documents
// @Produce json
// @Param id path string true "Document ID"
// @Success 200 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /documents/{id} [delete]
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
