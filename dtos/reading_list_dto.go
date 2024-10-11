package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReadingListInsertDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type ReadingListDTO struct {
	Id          primitive.ObjectID   `json:"_id,omitempty"`
	ReadingIds  []primitive.ObjectID `json:"reading_ids"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
}

type UpdateReadingsToListDTO struct {
	ReadingIds    []primitive.ObjectID `json:"reading_ids" validate:"required"`
	ReadingListId primitive.ObjectID   `json:"reading_list_id" validate:"required"`
}
