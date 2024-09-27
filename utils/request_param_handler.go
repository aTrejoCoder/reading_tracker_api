package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetObjectIdFromRequest(ctx *gin.Context) (primitive.ObjectID, error) {
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
