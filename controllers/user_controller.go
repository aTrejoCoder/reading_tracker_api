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

type UserController struct {
	userService services.UserService
	apiResponse utils.ApiResponse
	validator   *validator.Validate
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
		validator:   validator.New(),
	}
}

func (c UserController) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		userDTO, err := c.userService.GetUserById(userId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.NotFound(ctx, "User")
				return
			}

			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Found(ctx, userDTO, "user")
	}
}

func (c UserController) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userInsertDTO dtos.UserInsertDTO

		if err := ctx.ShouldBindJSON(&userInsertDTO); err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.validator.Struct(&userInsertDTO); err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		userDTO, err := c.userService.CreateUser(userInsertDTO)
		if err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Created(ctx, userDTO, "User")
	}
}

func (c UserController) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		var userInsertDTO dtos.UserInsertDTO

		if err := ctx.ShouldBindJSON(&userInsertDTO); err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.validator.Struct(&userInsertDTO); err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		userDTO, err := c.userService.UpdateUser(userId, userInsertDTO)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.NotFound(ctx, "User")
				return
			}

			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Updated(ctx, userDTO, "User")
	}
}

func (c UserController) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromRequest(ctx)
		if err != nil {
			c.apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = c.userService.DeleteUser(userId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.NotFound(ctx, "User")
				return
			}

			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.Deleted(ctx, "User")
	}
}
