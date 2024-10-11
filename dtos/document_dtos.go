package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomDocumentDTO struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	Description string             `bson:"description"`
	Content     string             `bson:"content"`
	FileURL     string             `bson:"file_url"`
	URL         string             `bson:"url"`
	Tags        []string           `bson:"tags"`
	Category    string             `bson:"category"`
	Version     string             `bson:"version"`
	Status      string             `bson:"status"`
}

type CustomDocumentInsertDTO struct {
	Title       string   `bson:"title" validate:"required"`
	Author      string   `bson:"author" validate:"required"`
	Description string   `bson:"description" validate:"omitempty,max=200"`
	Content     string   `bson:"content" validate:"omitempty"`
	FileURL     string   `bson:"file_url" validate:"omitempty,url"`
	URL         string   `bson:"url" validate:"omitempty,url"`
	Tags        []string `bson:"tags" validate:"omitempty,dive,max=50"`
	Category    string   `bson:"category" validate:"omitempty"`
	Version     string   `bson:"version" validate:"omitempty"`
	Status      string   `bson:"status" validate:"omitempty,oneof='draft' 'published' 'archived'"`
}
