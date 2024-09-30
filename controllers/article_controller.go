package controllers

import (
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/services"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ArticleController struct {
	articleService services.ArticleService
	apiResponse    utils.ApiResponse
	validator      *validator.Validate
}

func NewArticleController(articleService services.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: articleService,
		validator:      validator.New(),
	}
}

func (c ArticleController) GetArticleById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ArticleId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		ArticleDTO, err := c.articleService.GetArticleId(ArticleId)
		if err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Article")
			return
		}

		c.apiResponse.Found(ctx, ArticleDTO, "Article")
	}
}

func (c ArticleController) CreateArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ArticleInsertDTO dtos.ArticleInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &ArticleInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.articleService.CreateArticle(ArticleInsertDTO); err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, nil, "Article")
	}
}

func (c ArticleController) UpdateArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ArticleId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		var ArticleInsertDTO dtos.ArticleInsertDTO

		isJsonValidate := utils.BindAndValidate(ctx, &ArticleInsertDTO, c.validator, c.apiResponse)
		if !isJsonValidate {
			return
		}

		if err := c.articleService.UpdateArticle(ArticleId, ArticleInsertDTO); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Article")
			return
		}

		c.apiResponse.Updated(ctx, nil, "Article")
	}
}

func (c ArticleController) DeleteArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ArticleId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), 400)
			return
		}

		if err := c.articleService.DeleteArticle(ArticleId); err != nil {
			if !errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.ServerError(ctx, err.Error())
				return
			}

			c.apiResponse.NotFound(ctx, "Article")
			return
		}

		c.apiResponse.Deleted(ctx, "Article")
	}
}
