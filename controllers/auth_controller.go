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

		c.apiResponse.OK(ctx, jwtToken, "Signup succesfull")
	}
}

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
			c.apiResponse.ServerError(ctx, "can't proccess login")
			return
		}

		c.apiResponse.OK(ctx, jwtToken, "Login successful, welcome back "+userDTO.Username)
	}
}
