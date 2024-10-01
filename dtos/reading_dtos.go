package dtos

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingDTO struct {
	Id            primitive.ObjectID `json:"id"`
	ReadingType   string             `json:"reading_type"`
	ReadingStatus string             `json:"reading_status"`
	Notes         string             `json:"notes"`
	UpdatedAt     time.Time          `json:"last_update"`
}

type ReadingInsertDTO struct {
	ReadingType   string    `json:"reading_type" validate:"required,oneof=manga book document article"`
	UserId        string    `json:"client_id" validate:"required"`
	ReadingStatus string    `json:"reading_status" validate:"required,oneof=ongoing completed paused"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at" validate:"required"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ClientReadingInsertDTO struct {
	ReadingType     string                 `json:"reading_type"`
	ReadingsRecords []models.ReadingRecord `json:"reading_records"`
	ReadingStatus   string                 `json:"reading_status"`
	Notes           string                 `json:"notes"`
}
