package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(ctx *gin.Context, requestDTO interface{}, validator *validator.Validate, apiResponse ApiResponse) bool {

	if err := ctx.ShouldBindJSON(requestDTO); err != nil {
		apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
		return false
	}

	if err := validator.Struct(requestDTO); err != nil {
		apiResponse.Error(ctx, err.Error(), http.StatusBadRequest)
		return false
	}

	return true
}
