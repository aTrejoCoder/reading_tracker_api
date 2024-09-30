package repository

import (
	"context"
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository[T any] struct {
	collection *mongo.Collection
}

func NewRepository[T any](collection *mongo.Collection) *Repository[T] {
	return &Repository[T]{collection: collection}
}

func (r *Repository[T]) Create(ctx context.Context, document *T) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, document)
}

func (r *Repository[T]) GetByID(ctx context.Context, id primitive.ObjectID) (*T, error) {
	var result T
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *Repository[T]) Update(ctx context.Context, id primitive.ObjectID, update interface{}) (*mongo.UpdateResult, error) {
	updateDoc := bson.M{
		"$set": update,
	}
	mongoResult, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, updateDoc)
	if err != nil {
		return mongoResult, err
	}

	if mongoResult.MatchedCount == 0 {
		return mongoResult, utils.ErrNotFound
	}
	return mongoResult, nil
}

func (r *Repository[T]) DeleteById(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	mongoResult, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return mongoResult, err
	}

	if mongoResult.DeletedCount == 0 {
		return mongoResult, utils.ErrNotFound
	}
	return mongoResult, nil
}
