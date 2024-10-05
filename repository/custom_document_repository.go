package repository

import (
	"context"
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/dtos"
	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/aTrejoCoder/reading_tracker_api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomDocumentRepository struct {
	userCollection mongo.Collection
	userRepository Repository[models.User]
}

func NewCustomDocumentRepository(userCollection mongo.Collection, userRepository Repository[models.User]) *CustomDocumentRepository {
	return &CustomDocumentRepository{
		userCollection: userCollection,
		userRepository: userRepository,
	}
}

func (r CustomDocumentRepository) GetByUserId(ctx context.Context, userId primitive.ObjectID) ([]models.CustomDocument, error) {
	user, err := r.userRepository.GetByID(ctx, userId)
	if err != nil {
		return []models.CustomDocument{}, err
	}

	return user.CustomDocuments, nil
}

func (r CustomDocumentRepository) GetById(ctx context.Context, userId primitive.ObjectID, customDocumentId primitive.ObjectID) (*models.CustomDocument, error) {
	user, err := r.userRepository.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	customDocuments := user.CustomDocuments

	for _, customDocument := range customDocuments {
		if customDocument.Id == customDocumentId {
			return &customDocument, nil
		}
	}

	return nil, err
}

func (r CustomDocumentRepository) Create(ctx context.Context, userId primitive.ObjectID, customDocument models.CustomDocument) error {
	filter := bson.M{"_id": userId}

	update := bson.M{"$push": bson.M{
		"custom_documents": customDocument,
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

func (r CustomDocumentRepository) Update(ctx context.Context, userId primitive.ObjectID, customDocumentId primitive.ObjectID, customDocumentInsertDTO dtos.CustomDocumentInsertDTO) (*mongo.UpdateResult, error) {
	updateFields := bson.M{}

	// Check each field and only add it to the update map if it's not empty
	if customDocumentInsertDTO.Title != "" {
		updateFields["custom_documents.$[elem].title"] = customDocumentInsertDTO.Title
	}
	if customDocumentInsertDTO.Author != "" {
		updateFields["custom_documents.$[elem].author"] = customDocumentInsertDTO.Author
	}
	if customDocumentInsertDTO.Description != "" {
		updateFields["custom_documents.$[elem].description"] = customDocumentInsertDTO.Description
	}
	if customDocumentInsertDTO.FileURL != "" {
		updateFields["custom_documents.$[elem].file_url"] = customDocumentInsertDTO.FileURL
	}
	if customDocumentInsertDTO.URL != "" {
		updateFields["custom_documents.$[elem].url"] = customDocumentInsertDTO.URL
	}
	if len(customDocumentInsertDTO.Tags) > 0 {
		updateFields["custom_documents.$[elem].tags"] = customDocumentInsertDTO.Tags
	}
	if customDocumentInsertDTO.Category != "" {
		updateFields["custom_documents.$[elem].category"] = customDocumentInsertDTO.Category
	}
	if customDocumentInsertDTO.Version != "" {
		updateFields["custom_documents.$[elem].version"] = customDocumentInsertDTO.Version
	}
	if customDocumentInsertDTO.Status != "" {
		updateFields["custom_documents.$[elem].status"] = customDocumentInsertDTO.Status
	}
	// Always update the updated_at field
	updateFields["custom_documents.$[elem].updated_at"] = time.Now().UTC()

	// Filter to find the user by ID
	filter := bson.M{"_id": userId}

	// Array filter to match the specific custom document by its ID
	arrayFilter := bson.A{
		bson.M{"elem._id": customDocumentId}, // 'elem' refers to the array element
	}

	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: arrayFilter,
	})

	mongoResult, err := r.userCollection.UpdateOne(ctx, filter, bson.M{"$set": updateFields}, updateOptions)
	if err != nil {
		return nil, err
	}

	if mongoResult.ModifiedCount == 0 {
		return nil, utils.ErrNotFound
	}

	return mongoResult, nil
}

func (r CustomDocumentRepository) Delete(ctx context.Context, userId primitive.ObjectID, customDocumentId primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": userId}

	update := bson.M{
		"$pull": bson.M{
			"custom_documents": bson.M{"_id": customDocumentId},
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
