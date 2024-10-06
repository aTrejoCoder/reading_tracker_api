package repository

import (
	"context"
	"log"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
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

func (r *ReadingExtendRepository) UpdateRecord(ctx context.Context, id primitive.ObjectID, userId primitive.ObjectID, update interface{}) (*mongo.UpdateResult, error) {
	log.Println("Starting UpdateRecord function")
	log.Printf("Updating record with ID: %s and UserID: %s\n", id.Hex(), userId.Hex())

	filter := bson.M{
		"_id":     id,
		"user_id": userId,
	}
	log.Printf("Filter: %+v\n", filter)
	log.Printf("Update data: %+v\n", update)

	mongoResult, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error in UpdateOne: %v\n", err)
		return mongoResult, err
	}

	if mongoResult.MatchedCount == 0 {
		log.Println("No documents matched the filter")
		return mongoResult, utils.ErrNotFound
	}

	log.Printf("Update result: %+v\n", mongoResult)
	return mongoResult, nil
}

func (r *ReadingExtendRepository) UpdateWithFilter(ctx context.Context, id primitive.ObjectID, update interface{}, arrayFilter interface{}) (*mongo.UpdateResult, error) {
	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: arrayFilter.(primitive.A),
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

func (r *ReadingExtendRepository) UpdateRecordProperties(ctx context.Context, readingId primitive.ObjectID, recordId primitive.ObjectID, recordUpdateDTO dtos.ReadingRecordInsertDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id": readingId,
	}

	update := bson.M{
		"$set": bson.M{
			"reading_records.$[record].notes":    recordUpdateDTO.Notes,
			"reading_records.$[record].progress": recordUpdateDTO.Progress,
		},
	}

	arrayFilter := bson.A{
		bson.M{"record._id": recordId},
	}

	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: arrayFilter,
	})

	mongoResult, err := r.collection.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return mongoResult, nil
}

func (r *ReadingExtendRepository) GetReadingsByUserId(ctx context.Context, userId primitive.ObjectID, page, limit int64, sortOrder int) ([]models.Reading, error) {
	filter := bson.M{"user_id": userId}

	skip := (page - 1) * limit
	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: "created_at", Value: sortOrder}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
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

func (r *ReadingExtendRepository) GetReadingsByUserIdAndReadingType(
	ctx context.Context,
	userId primitive.ObjectID,
	readingType string,
	sortValue string, // created_at, updated_at, last_record_update
	page, limit int64,
	sortOrder int) ([]models.Reading, error) {

	filter := bson.M{"user_id": userId, "reading_type": readingType}

	skip := (page - 1) * limit
	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: sortValue, Value: sortOrder}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
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

func (r *ReadingExtendRepository) GetReadingsByUserIdAndReadingStatus(
	ctx context.Context,
	userId primitive.ObjectID,
	readingType string,
	page, limit int64,
	sortOrder int) ([]models.Reading, error) {

	filter := bson.M{"user_id": userId, "reading_status": readingType}

	skip := (page - 1) * limit
	findOptions := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.D{{Key: "updated_at", Value: sortOrder}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
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

func (r *ReadingExtendRepository) IsReadingExistsForUser(ctx context.Context, userId, documentId primitive.ObjectID) (bool, error) {
	filter := bson.M{
		"user_id":     userId,
		"document_id": documentId,
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
