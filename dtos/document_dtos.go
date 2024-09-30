package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type DocumentDTO struct {
	Id          primitive.ObjectID `json:"_id,omitempty"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Description string             `json:"description"`
	FileURL     string             `json:"file_url"`
	FileType    string             `json:"file_type"`
	Tags        []string           `json:"tags"`
	Version     string             `json:"version"`
	Status      string             `json:"status"`
}

type DocumentInsertDTO struct {
	Title       string   `json:"title" validate:"required"`
	Author      string   `json:"author" validate:"required"`
	Description string   `json:"description" validate:"required"`
	FileURL     string   `json:"file_url"`
	FileType    string   `json:"file_type" validate:"required"`
	Tags        []string `json:"tags"`
	Version     string   `json:"version"`
	Status      string   `json:"status"`
}
