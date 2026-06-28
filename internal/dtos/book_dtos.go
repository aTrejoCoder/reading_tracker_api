package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookDTO struct {
	Id              primitive.ObjectID `json:"_id,omitempty"`
	Author          string             `json:"author"`
	ISBN            string             `json:"ISBN"`
	Name            string             `json:"name"`
	CoverImageURL   string             `json:"cover_image_url"`
	Edition         string             `json:"edition"`
	Pages           int                `json:"pages"`
	Language        string             `json:"language"`
	PublicationDate time.Time          `json:"publication_date"`
	Publisher       string             `json:"publisher"`
	Description     string             `json:"description"`
	Genres          []string           `json:"genres"`
}

type BookInsertDTO struct {
	Author          string    `json:"author" validate:"required"`
	ISBN            string    `json:"ISBN" `
	Name            string    `json:"name" validate:"required"`
	CoverImageURL   string    `json:"cover_image_url"`
	Edition         string    `json:"edition"`
	Pages           int       `json:"pages" validate:"required"`
	Language        string    `json:"language" validate:"required"`
	PublicationDate time.Time `json:"publication_date"`
	Publisher       string    `json:"publisher"`
	Description     string    `json:"description"`
	Genres          []string  `json:"genres"`
}
