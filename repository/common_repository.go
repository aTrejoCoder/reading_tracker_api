package repository

import (
	"context"
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r *Repository[T]) GetByFilter(ctx context.Context, filter interface{}) (*T, error) {
	var result T
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (r *Repository[T]) GetManyByFilter(ctx context.Context, filter interface{}) ([]T, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	for cursor.Next(ctx) {
		var elem T
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
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

func (r *Repository[T]) GetManyByFilterPaginated(ctx context.Context, filter interface{}, page int64, limit int64) ([]T, error) {
	sortFields := bson.D{
		{Key: "name", Value: 1},
		{Key: "author", Value: 1},
	}

	opts := options.Find().
		SetSort(sortFields).
		SetLimit(limit).
		SetSkip((page - 1) * limit)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	for cursor.Next(ctx) {
		var elem T
		err := cursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, utils.ErrNotFound
	}

	return results, nil
}

func (r *Repository[T]) GetAllSortedPaginated(ctx context.Context, sortFields bson.D, page int64, limit int64) ([]T, error) {
	opts := options.Find().
		SetSort(sortFields).
		SetLimit(limit).
		SetSkip((page - 1) * limit)

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	for cursor.Next(ctx) {
		var elem T
		if err := cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
