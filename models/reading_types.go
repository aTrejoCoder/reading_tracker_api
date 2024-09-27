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

type Article struct {
	Id            primitive.ObjectID   `bson:"_id,omitempty"`
	Title         string               `bson:"title"`
	Author        string               `bson:"author"`
	Content       string               `bson:"content"`
	Summary       string               `bson:"summary"`
	PublishedDate time.Time            `bson:"published_date"`
	Tags          []string             `bson:"tags"`
	Category      string               `bson:"category"`
	URL           string               `bson:"url"`
	CreatedAt     time.Time            `bson:"created_at"`
	UpdatedAt     time.Time            `bson:"updated_at"`
	ReadingList   []primitive.ObjectID `bson:"reading_list"`
	Status        string               `bson:"status"`
	Views         int                  `bson:"views"`
}

type Document struct {
	Id          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"title"`
	Author      string               `bson:"author"`
	Description string               `bson:"description"`
	FileURL     string               `bson:"file_url"`
	FileType    string               `bson:"file_type"`
	Tags        []string             `bson:"tags"`
	CreatedAt   time.Time            `bson:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at"`
	Version     string               `bson:"version"`
	ReadingList []primitive.ObjectID `bson:"reading_list"`
	Status      string               `bson:"status"`
}

type Manga struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty"`
	Title           string               `bson:"title"`
	Author          string               `bson:"author"`
	CoverImageURL   string               `bson:"cover_image_url"`
	Volume          int                  `bson:"volume"`
	Chapters        int                  `bson:"chapters"`
	Demogragphy     string               `bson:"demograhpy"`
	Genres          []string             `bson:"genres"`
	PublicationDate time.Time            `bson:"publication_date"`
	Publisher       string               `bson:"publisher"`
	Description     string               `bson:"description"`
	ReadingList     []primitive.ObjectID `bson:"reading_list"`
	CreatedAt       time.Time            `bson:"created_at"`
	UpdatedAt       time.Time            `bson:"updated_at"`
}
