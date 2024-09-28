package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDTO struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Profile  ProfileDTO         `bson:"profile"`
	Roles    []string           `bson:"roles"`
}

type UserInsertDTO struct {
	Username string `bson:"username" validate:"required,min=3,max=32"`
	Email    string `bson:"email" validate:"required,email"`
	Password string `bson:"password" validate:"required,min=8"`
}

type ProfileDTO struct {
	FullName        string `bson:"full_name"`
	Biography       string `bson:"biography"`
	ProfileImageURL string `bson:"profile_image_url"`
	ProfileCoverURL string `bson:"profile_cover_url"`
}

type ProfileInsertDTO struct {
	FirstName       string `json:"first_name"  validate:"required,min=3"`
	LastName        string `json:"last_name" validate:"required,min=3"`
	Biography       string `json:"biography"`
	ProfileImageURL string `json:"profile_image_url"`
	ProfileCoverURL string `json:"profile_cover_url"`
}
