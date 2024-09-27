package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	Profile      Profile            `bson:"profile"`
	LastLogin    time.Time          `bson:"last_login"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	ReadingsList ReadingsList       `bson:"readings"`
	Roles        []string           `bson:"roles"`
}

type Profile struct {
	FullName        string    `bson:"full_name"`
	Biography       string    `bson:"biography"`
	ProfileImageURL string    `bson:"profile_image_url"`
	ProfileCoverURL string    `bson:"profile_cover_url"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}
