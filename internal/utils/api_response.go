package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ApiResponse represents the structure of an API response.
// @Description Represents a standardized API response structure.
// @Property success boolean Indicates if the request was successful
// @Property code integer The HTTP status code
// @Property data object The data returned by the API
// @Property msg string A message providing additional context about the response
// @Property timestamp string The time at which the response was generated
type ApiResponse struct {
	Success   bool        `json:"success"`   // Indicates if the request was successful
	Code      int         `json:"code"`      // The HTTP status code
	Data      interface{} `json:"data"`      // The data returned by the API
	Msg       string      `json:"msg"`       // A message providing additional context
	Timestamp time.Time   `json:"timestamp"` // The time at which the response was generated
}

// Error returns an error response.
// @Success 400 {object} ApiResponse
// @Failure 500 {object} ApiResponse
func (r ApiResponse) Error(ctx *gin.Context, msg string, code int) {
	ctx.JSON(code, ApiResponse{
		Success:   false,
		Code:      code,
		Data:      nil,
		Msg:       msg,
		Timestamp: time.Now(),
	})
}

// NotFound returns a not found response.
// @Success 404 {object} ApiResponse
// @Router /path-to-resource [get]
func (r ApiResponse) NotFound(ctx *gin.Context, entity string) {
	ctx.JSON(http.StatusNotFound, ApiResponse{
		Success:   false,
		Code:      http.StatusNotFound,
		Data:      nil,
		Msg:       entity + " Not Found",
		Timestamp: time.Now(),
	})
}

// ServerError returns a server error response.
// @Failure 500 {object} ApiResponse
// @Router /path-to-resource [post]
func (r ApiResponse) ServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, ApiResponse{
		Success:   false,
		Code:      http.StatusInternalServerError,
		Data:      nil,
		Msg:       msg,
		Timestamp: time.Now(),
	})
}

// OK returns a successful response.
// @Success 200 {object} ApiResponse
// @Router /path-to-resource [get]
func (r ApiResponse) OK(ctx *gin.Context, data any, msg string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      data,
		Msg:       msg,
		Timestamp: time.Now(),
	})
}

// Found returns a found response.
// @Success 200 {object} ApiResponse
// @Router /path-to-resource [get]
func (r ApiResponse) Found(ctx *gin.Context, data any, entity string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      data,
		Msg:       entity + " Successfully Retrieved",
		Timestamp: time.Now(),
	})
}

// Created returns a created response.
// @Success 201 {object} ApiResponse
// @Router /path-to-resource [post]
func (r ApiResponse) Created(ctx *gin.Context, data any, entity string) {
	ctx.JSON(http.StatusCreated, ApiResponse{
		Success:   true,
		Code:      http.StatusCreated,
		Data:      data,
		Msg:       entity + " Successfully Created",
		Timestamp: time.Now(),
	})
}

// Updated returns an updated response.
// @Success 200 {object} ApiResponse
// @Router /path-to-resource [put]
func (r ApiResponse) Updated(ctx *gin.Context, data any, entity string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      data,
		Msg:       entity + " Successfully Updated",
		Timestamp: time.Now(),
	})
}

// Deleted returns a deleted response.
// @Success 200 {object} ApiResponse
// @Router /path-to-resource [delete]
func (r ApiResponse) Deleted(ctx *gin.Context, entity string) {
	ctx.JSON(http.StatusOK, ApiResponse{
		Success:   true,
		Code:      http.StatusOK,
		Data:      nil,
		Msg:       entity + " Successfully Deleted",
		Timestamp: time.Now(),
	})
}
