package repository

import (
	"context"
	"errors"
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReadingListRepository struct {
	userCollection mongo.Collection
	userRepository Repository[models.User]
}

func NewReadingListRepository(userCollection mongo.Collection, userRepository Repository[models.User]) *ReadingListRepository {
	return &ReadingListRepository{
		userCollection: userCollection,
		userRepository: userRepository,
	}
}

func (r ReadingListRepository) GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]models.ReadingsList, error) {
	user, err := r.userRepository.GetByID(ctx, userId)
	if err != nil {
		return []models.ReadingsList{}, err
	}

	return user.ReadingsLists, nil
}

func (r ReadingListRepository) CreateReadingList(ctx context.Context, userId primitive.ObjectID, readingList models.ReadingsList) error {
	filter := bson.M{"_id": userId}

	update := bson.M{"$push": bson.M{
		"reading_lists": readingList,
	}}

	mongoResult, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if mongoResult.ModifiedCount == 0 {
		return utils.ErrNotFound
	}

	return nil
}

func (r ReadingListRepository) AddReadingsToList(ctx context.Context, userId primitive.ObjectID, readingListId primitive.ObjectID, readingsId []primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id":               userId,
		"reading_lists._id": readingListId,
	}

	update := bson.M{
		"$addToSet": bson.M{
			"reading_lists.$.reading_ids": bson.M{
				"$each": readingsId,
			},
		},
	}

	mongoResult, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if mongoResult.MatchedCount == 0 {
		return nil, utils.ErrNotFound
	}

	if mongoResult.ModifiedCount == 0 {
		return mongoResult, nil
	}

	return mongoResult, nil
}

func (r ReadingListRepository) RemoveReadingsFromList(ctx context.Context, userId primitive.ObjectID, readingListId primitive.ObjectID, readingsId []primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"_id":               userId,
		"reading_lists._id": readingListId,
	}

	update := bson.M{
		"$pull": bson.M{
			"reading_lists.$.reading_ids": bson.M{"$in": readingsId},
		},
	}

	mongoResult, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if mongoResult.MatchedCount == 0 {
		return nil, err
	}

	if mongoResult.ModifiedCount == 0 {
		return nil, errors.New("no changes")
	}
	return mongoResult, nil
}

func (r ReadingListRepository) UpdateReadingList(ctx context.Context, userId primitive.ObjectID, readingListId primitive.ObjectID, insertDTO dtos.ReadingListInsertDTO) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": userId}

	update := bson.M{
		"$set": bson.M{
			"reading_lists.$[list].name":        insertDTO.Name,
			"reading_lists.$[list].description": insertDTO.Description,
			"reading_lists.$[list].updated_at":  time.Now().UTC(),
		},
	}

	arrayFilter := bson.A{
		bson.M{"list._id": readingListId},
	}

	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: arrayFilter,
	})

	mongoResult, err := r.userCollection.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}

	if mongoResult.ModifiedCount == 0 {
		return nil, utils.ErrNotFound
	}

	return mongoResult, nil
}

func (r ReadingListRepository) DeleteReadingList(ctx context.Context, userId primitive.ObjectID, readingListId primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": userId}

	update := bson.M{
		"$pull": bson.M{
			"reading_lists": bson.M{"_id": readingListId},
		},
	}

	mongoResult, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if mongoResult.ModifiedCount == 0 {
		return nil, utils.ErrNotFound
	}

	return mongoResult, nil
}
