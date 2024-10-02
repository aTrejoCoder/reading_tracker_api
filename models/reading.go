package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reading struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	UserId          primitive.ObjectID `bson:"user_id"`
	ReadingType     string             `bson:"reading_type"`
	ReadingsRecords []ReadingRecord    `bson:"reading_records"`
	ReadingStatus   string             `bson:"reading_status"`
	Notes           string             `bson:"notes"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type ReadingRecord struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Progress   string             `bson:"progress"`
	Notes      string             `bson:"notes"`
	RecordDate time.Time          `bson:"update_date"`
}

type ReadingsList struct {
	BookReadings     []Reading `bson:"book_readings"`
	MangaReadings    []Reading `bson:"manga_readings"`
	DocumentReadings []Reading `bson:"document_readings"`
	ArticleReadings  []Reading `bson:"article_readings"`
}
