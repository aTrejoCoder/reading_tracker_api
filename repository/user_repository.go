package repository

import (
	"context"
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateLastLogin(ctx context.Context, userId primitive.ObjectID) error
}

type userRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepositoryImpl{
		collection: collection,
	}
}

func (ur userRepositoryImpl) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := ur.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Not found
			return nil, utils.ErrNotFound
		} else {
			return nil, utils.ErrDatabase
		}
	}
	return &user, nil
}

func (ur userRepositoryImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := ur.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Not found
			return nil, utils.ErrNotFound
		} else {
			return nil, utils.ErrDatabase
		}
	}
	return &user, nil
}

func (ur userRepositoryImpl) UpdateLastLogin(ctx context.Context, userId primitive.ObjectID) error {
	filter := bson.M{"_id": userId}

	update := bson.M{
		"$set": bson.M{
			"last_login": time.Now().UTC(),
		},
	}

	_, err := ur.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
