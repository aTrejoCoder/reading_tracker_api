package utils

import (
	"errors"
	"fmt"
	"net/http"
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

	userObjectId, _ := primitive.ObjectIDFromHex(userIdStr)

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
