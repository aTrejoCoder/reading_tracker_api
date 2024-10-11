package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MangaDTO struct {
	Id              primitive.ObjectID `json:"_id,omitempty"`
	Title           string             `json:"title"`
	Author          string             `json:"author"`
	CoverImageURL   string             `json:"cover_image_url"`
	Volume          int                `json:"volume"`
	Chapters        int                `json:"chapters"`
	Demogragphy     string             `json:"demograhpy"`
	Genres          []string           `json:"genres"`
	PublicationDate time.Time          `json:"publication_date"`
	Publisher       string             `json:"publisher"`
	Description     string             `json:"description"`
}

type MangaInsertDTO struct {
	Title           string    `json:"title" validate:"required"`
	Author          string    `json:"author" validate:"required"`
	CoverImageURL   string    `json:"cover_image_url"`
	Volume          int       `json:"volume" validate:"required"`
	Chapters        int       `json:"chapters" validate:"required"`
	Demogragphy     string    `json:"demograhpy" validate:"required"`
	Genres          []string  `json:"genres" validate:"required"`
	PublicationDate time.Time `json:"publication_date"`
	Publisher       string    `json:"publisher"`
	Description     string    `json:"description"`
}
