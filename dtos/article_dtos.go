package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleDTO struct {
	Id            primitive.ObjectID `json:"_id,omitempty"`
	Title         string             `json:"title"`
	Author        string             `json:"author"`
	Content       string             `json:"content"`
	Summary       string             `json:"summary"`
	PublishedDate time.Time          `json:"published_date"`
	Tags          []string           `json:"tags"`
	Category      string             `json:"category"`
	URL           string             `json:"url"`
	Status        string             `json:"status"`
}

type ArticleInsertDTO struct {
	Title         string    `json:"title" validate:"required,min=5,max=100"`
	Author        string    `json:"author" validate:"required,min=3,max=50"`
	Content       string    `json:"content" validate:"required,min=20"`
	Summary       string    `json:"summary" validate:"required,max=255"`
	PublishedDate time.Time `json:"published_date" validate:"required"`
	Tags          []string  `json:"tags" validate:"required,dive,required"`
	Category      string    `json:"category" validate:"required,oneof=Technology Science Health Business"`
	URL           string    `json:"url" validate:"required,url"`
	Status        string    `json:"status" validate:"required,oneof=published draft archived"`
}
