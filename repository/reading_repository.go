package repository

import (
	"context"
	"errors"

	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReadingExtendRepository struct {
	collection mongo.Collection
}

func NewReadingExtendRepository(collection mongo.Collection) *ReadingExtendRepository {
	return &ReadingExtendRepository{collection: collection}
}

func (r *ReadingExtendRepository) UpdateRecord(ctx context.Context, id primitive.ObjectID, update interface{}) (*mongo.UpdateResult, error) {
	mongoResult, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return mongoResult, err
	}

	if mongoResult.MatchedCount == 0 {
		return mongoResult, utils.ErrNotFound
	}
	return mongoResult, nil
}

func (r *ReadingExtendRepository) UpdateWithFilter(ctx context.Context, id primitive.ObjectID, update interface{}, arrayFilter interface{}) (*mongo.UpdateResult, error) {
	var filters []interface{}

	switch af := arrayFilter.(type) {
	case primitive.A:
		filters = af
	case []interface{}:
		filters = af
	default:
		return nil, errors.New("arrayFilter debe ser de tipo []interface{} o primitive.A")
	}

	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: filters,
	})

	mongoResult, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update, updateOptions)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return mongoResult, utils.ErrNotFound
		}
		return mongoResult, err
	}
	return mongoResult, nil
}
func (r *ReadingExtendRepository) GetReadingsByUserId(ctx context.Context, userId primitive.ObjectID) ([]models.Reading, error) {
	filter := bson.M{"user_id": userId}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var readings []models.Reading

	for cursor.Next(ctx) {
		var reading models.Reading
		if err := cursor.Decode(&reading); err != nil {
			return nil, err
		}
		readings = append(readings, reading)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return readings, nil
}
