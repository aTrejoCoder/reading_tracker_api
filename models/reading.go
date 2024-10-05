package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reading struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	UserId          primitive.ObjectID `bson:"user_id"`
	DocumentId      primitive.ObjectID `bson:"user_id"`
	ReadingType     string             `bson:"reading_type"` // Book, Manga, CustomDoc
	ReadingsRecords []ReadingRecord    `bson:"reading_records"`
	ReadingStatus   string             `bson:"reading_status"`
	Notes           string             `bson:"notes"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type ReadingRecord struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Notes      string             `bson:"notes"`
	RecordDate time.Time          `bson:"update_date"`
}

type ReadingsList struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"`
	ReadingIds  []primitive.ObjectID `bson:"reading_ids"`
	Name        string               `bson:"name"`
	Description string               `bson:"description"`
	CreatedAt   time.Time            `bson:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at"`
}
