package dtos

import (
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadingDTO struct {
	Id            primitive.ObjectID `json:"id"`
	UserId        string             `json:"client_id"`
	ReadingType   string             `json:"reading_type"`
	ReadingStatus string             `json:"reading_status"`
	Notes         string             `json:"notes"`
	UpdatedAt     time.Time          `json:"last_update"`
	RecordsDTOs   []ReadingRecordDTO `json:"records"`
}

type ReadingInsertDTO struct {
	ReadingType   string    `json:"reading_type" validate:"required,oneof=manga book document article"`
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

type ReadingRecordInsertDTO struct {
	ReadingId primitive.ObjectID `json:"reading_id" validate:"required"`
	Progress  string             `json:"progress" validate:"required"`
	Notes     string             `json:"notes" validate:"required"`
}

type ReadingRecordDTO struct {
	Id         primitive.ObjectID `json:"id"`
	Progress   string             `json:"progress"`
	Notes      string             `json:"notes"`
	RecordDate time.Time          `bson:"update_date"`
}
