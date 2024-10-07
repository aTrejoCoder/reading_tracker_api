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

type AuthController struct {
	authServices services.AuthServices
	apiResponse  utils.ApiResponse
	validator    *validator.Validate
}

func NewAuthController(authServices services.AuthServices) *AuthController {
	return &AuthController{
		authServices: authServices,
		validator:    validator.New(),
	}
}

// Signup handles user registration
// @Summary User Signup
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param signup body dtos.SignupDTO true "User Signup Information"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 409 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /auth/signup [post]
func (c AuthController) Signup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signUpDTO dtos.SignupDTO

		if !utils.BindAndValidate(ctx, &signUpDTO, c.validator, c.apiResponse) {
			return
		}

		if err := c.authServices.ValidateSignupCredentials(signUpDTO); err != nil {
			if !errors.Is(err, utils.ErrDatabase) {
				c.apiResponse.Error(ctx, err.Error(), http.StatusConflict)
				return
			}
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		jwtToken, err := c.authServices.ProccesSignup(signUpDTO)
		if err != nil {
			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		c.apiResponse.OK(ctx, jwtToken, "Signup successful")
	}
}

// Login handles user authentication
// @Summary User Login
// @Description Authenticate an existing user
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dtos.LoginDTO true "User Login Information"
// @Success 200 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Failure 401 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /auth/login [post]
func (c AuthController) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginDTO dtos.LoginDTO

		if !utils.BindAndValidate(ctx, &loginDTO, c.validator, c.apiResponse) {
			return
		}

		userDTO, err := c.authServices.ValidateLoginCredentials(loginDTO)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				c.apiResponse.NotFound(ctx, "User with given credentials")
				return
			}

			if errors.Is(err, utils.ErrUnauthorized) {
				c.apiResponse.Error(ctx, err.Error(), http.StatusUnauthorized)
				return
			}

			c.apiResponse.ServerError(ctx, err.Error())
			return
		}

		jwtToken, err := c.authServices.ProccesLogin(*userDTO)
		if err != nil {
			c.apiResponse.ServerError(ctx, "can't process login")
			return
		}

		c.apiResponse.OK(ctx, jwtToken, "Login successful, welcome back "+userDTO.Username)
	}
}
