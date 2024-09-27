package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	Timestamp time.Time   `json:"timestamp"`
}

func (r ApiResponse) Error(ctx *gin.Context, msg string, code int) {
	ctx.JSON(code, ApiResponse{
		Success:   false,
		Code:      code,
		Data:      nil,
		Msg:       msg,
		Timestamp: time.Now(),
	})
}

func (r ApiResponse) NotFound(ctx *gin.Context, entity string) {
	ctx.JSON(http.StatusNotFound, ApiResponse{
		Success:   false,
		Code:      http.StatusNotFound,
		Data:      nil,
		Msg:       entity + " Not Found",
		Timestamp: time.Now(),
	})
}

func (r ApiResponse) ServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, ApiResponse{
		Success:   false,
		Code:      http.StatusInternalServerError,
		Data:      nil,
		Msg:       msg,
		Timestamp: time.Now(),
	})
}

func (r ApiResponse) Found(ctx *gin.Context, data any, entity string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      data,
		Msg:       entity + " Successfully Retrieved",
		Timestamp: time.Now(),
	})
}

func (r ApiResponse) Created(ctx *gin.Context, data any, entity string) {
	ctx.JSON(http.StatusCreated, ApiResponse{
		Success:   true,
		Code:      http.StatusCreated,
		Data:      data,
		Msg:       entity + " Successfully Created",
		Timestamp: time.Now(),
	})
}

func (r ApiResponse) Updated(ctx *gin.Context, data any, entity string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      data,
		Msg:       entity + " Successfully Updated",
		Timestamp: time.Now(),
	})
}

func (r ApiResponse) Deleted(ctx *gin.Context, entity string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      nil,
		Msg:       entity + " Successfully Deleted",
		Timestamp: time.Now(),
	})
}
