package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty"`
	Author          string               `bson:"author"`
	ISBN            string               `bson:"ISBN"`
	Name            string               `bson:"name"`
	CoverImageURL   string               `bson:"cover_image_url"`
	Edition         string               `bson:"edition"`
	Pages           int                  `bson:"pages"`
	Language        string               `bson:"language"`
	PublicationDate time.Time            `bson:"publication_date"`
	Publisher       string               `bson:"publisher"`
	Description     string               `bson:"description"`
	Genres          []string             `bson:"genres"`
	ReadingList     []primitive.ObjectID `bson:"reading_list"`
	CreatedAt       time.Time            `bson:"created_at"`
	UpdatedAt       time.Time            `bson:"updated_at"`
}

type Manga struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title"`
	Author          string             `bson:"author"`
	CoverImageURL   string             `bson:"cover_image_url"`
	Volume          int                `bson:"volume"`
	Chapters        int                `bson:"chapters"`
	Demography      string             `bson:"demograhpy"`
	Genres          []string           `bson:"genres"`
	PublicationDate time.Time          `bson:"publication_date"`
	Publisher       string             `bson:"publisher"`
	Description     string             `bson:"description"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type CustomDocument struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	Description string             `bson:"description"`
	Content     string             `bson:"content"`
	FileURL     string             `bson:"file_url"`
	URL         string             `bson:"url"`
	Tags        []string           `bson:"tags"`
	Category    string             `bson:"category"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Version     string             `bson:"version"`
	Status      string             `bson:"status"`
}
