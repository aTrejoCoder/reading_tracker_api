package test

import (
	"context"
	"testing"
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestReadingInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	readingCollection := getCollection("readings_test")

	defer closeTestDB()
	defer dropCollection(readingCollection)

	readingTest := getTestReading()

	// Insert
	_, err := readingCollection.InsertOne(context.TODO(), readingTest)
	if err != nil {
		t.Fatalf("Failed to insert reading: %v", err)
	}

	// Retrieve the reading by ID (_id)
	var retrievedReading models.Reading
	err = readingCollection.FindOne(context.TODO(), bson.M{"_id": readingTest.Id}).Decode(&retrievedReading)
	if err != nil {
		t.Fatalf("Failed to retrieve reading: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, readingTest.Id, retrievedReading.Id)
	assert.Equal(t, readingTest.ReadingType, retrievedReading.ReadingType)
	assert.Equal(t, readingTest.ReadingStatus, retrievedReading.ReadingStatus)
	assert.Equal(t, readingTest.CreatedAt.Truncate(time.Millisecond), retrievedReading.CreatedAt.Truncate(time.Millisecond))
	assert.Equal(t, readingTest.UpdatedAt.Truncate(time.Millisecond), retrievedReading.UpdatedAt.Truncate(time.Millisecond))
}

func TestReadingUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	readingCollection := getCollection("readings_test")

	defer closeTestDB()
	defer dropCollection(readingCollection)

	readingTest := getTestReading()

	// Create
	_, err := readingCollection.InsertOne(context.TODO(), readingTest)
	if err != nil {
		t.Fatalf("Failed to insert reading: %v", err)
	}

	// Update
	filter := bson.M{"_id": readingTest.Id}
	update := bson.M{
		"$set": bson.M{
			"reading_type":  "manga",
			"reading_track": 1,
			"updated_at":    time.Now().UTC(),
		},
	}
	_, err = readingCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to update reading: %v", err)
	}

	// Retrieve the updated reading
	var updatedReading models.Reading
	err = readingCollection.FindOne(context.TODO(), bson.M{"_id": readingTest.Id}).Decode(&updatedReading)
	if err != nil {
		t.Fatalf("Failed to retrieve updated reading: %v", err)
	}

	// Compare updated fields
	assert.NoError(t, err)
	assert.Equal(t, readingTest.Id, updatedReading.Id)
	assert.Equal(t, "manga", updatedReading.ReadingType)
	assert.Equal(t, "PLAN_TO_READ", updatedReading.ReadingStatus)
	assert.NotEqual(t, readingTest.UpdatedAt, updatedReading.UpdatedAt)
}

func TestReadingDeletion(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	readingCollection := getCollection("readings_test")

	defer closeTestDB()
	defer dropCollection(readingCollection)

	// Insert
	readingTest := getTestReading()
	_, err := readingCollection.InsertOne(context.TODO(), readingTest)
	if err != nil {
		t.Fatalf("Failed to insert reading: %v", err)
	}

	// Delete
	filter := bson.M{"_id": readingTest.Id}
	_, err = readingCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Fatalf("Failed to delete reading: %v", err)
	}

	// Verify Delete
	var deletedReading models.Reading
	err = readingCollection.FindOne(context.TODO(), filter).Decode(&deletedReading)

	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}

func getTestReading() models.Reading {
	now := time.Now().UTC()
	return models.Reading{
		Id:              primitive.NewObjectID(),
		UserId:          primitive.NewObjectID(),
		ReadingType:     "book",
		ReadingStatus:   "PLAN_TO_READ",
		ReadingsRecords: []models.ReadingRecord{},
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
