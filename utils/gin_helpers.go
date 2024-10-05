package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aTrejoCoder/reading_tracker_api/middleware/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetUserIdFromRequest(ctx *gin.Context, apiResponse ApiResponse) (primitive.ObjectID, bool) {
	jwtToken, err := extractJWT(ctx)
	if err != nil {
		apiResponse.Error(ctx, err.Error(), http.StatusUnauthorized)
		return primitive.ObjectID{}, false
	}

	userIdStr, err := token.GetIDFromJWT(jwtToken)
	if err != nil {
		apiResponse.Error(ctx, ErrUnauthorized.Error(), http.StatusUnauthorized)
		return primitive.ObjectID{}, false
	}

	userObjectId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		apiResponse.Error(ctx, userIdStr, http.StatusInternalServerError)
		return primitive.ObjectID{}, false
	}

	return userObjectId, true
}

func GetObjectIdFromUrlParam(ctx *gin.Context) (primitive.ObjectID, error) {
	id := ctx.Param("id")

	if id == "" {
		return primitive.ObjectID{}, errors.New("id not provided")
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid objectId")
	}

	return objectId, nil
}

func GetObjectIdFromUrlQuery(ctx *gin.Context, id string) (primitive.ObjectID, error) {
	queryId := ctx.Query(id)

	if queryId == "" {
		return primitive.ObjectID{}, errors.New("query id not provided")
	}

	objectId, err := primitive.ObjectIDFromHex(queryId)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid objectId")
	}

	return objectId, nil
}

func extractJWT(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no Authorization header provided")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	return parts[1], nil
}

func GetPaginationValuesFromRequest(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	var page, limit int64

	if pageStr == "" {
		page = 1 // Default Page
	} else {
		var err error
		page, err = strconv.ParseInt(pageStr, 10, 64)
		if err != nil || page < 1 {
			page = 1 // Default Page
		}
	}

	if limitStr == "" {
		limit = 10 // Default Limit
	} else {
		var err error
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil || limit < 1 {
			limit = 10 // Default Limit
		}
	}

	return page, limit
}
