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

// UserController handles user-related operations.
type UserController struct {
	userService services.UserService
	apiResponse utils.ApiResponse
	validator   *validator.Validate
}

// NewUserController creates a new UserController.
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
		validator:   validator.New(),
	}
}

// GetUserById retrieves a user by ID.
// @Summary Get User by ID
// @Description Retrieve a user by their ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} dtos.UserDTO "Success"
// @Failure 400 {object} utils.ApiResponse "Invalid User ID"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/users/{userId} [get]
func (c UserController) GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlParam(ctx)
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

// CreateUser creates a new user.
// @Summary Create User
// @Description Create a new user in the system.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dtos.UserInsertDTO true "User data"
// @Success 201 {object} dtos.UserDTO "User created"
// @Failure 400 {object} utils.ApiResponse "Validation error"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/users [post]
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

// UpdateUser updates an existing user.
// @Summary Update User
// @Description Update an existing user in the system.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param user body dtos.UserInsertDTO true "Updated user data"
// @Success 200 {object} dtos.UserDTO "User updated"
// @Failure 400 {object} utils.ApiResponse "Validation error"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/users/{userId} [put]
func (c UserController) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlParam(ctx)
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

// DeleteUser deletes a user by ID.
// @Summary Delete User
// @Description Delete a user from the system by their ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 204 {object} utils.ApiResponse "User deleted"
// @Failure 404 {object} utils.ApiResponse "User not found"
// @Failure 500 {object} utils.ApiResponse "Internal Server Error"
// @Router /api/users/{userId} [delete]
func (c UserController) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, err := utils.GetObjectIdFromUrlParam(ctx)
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
