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
