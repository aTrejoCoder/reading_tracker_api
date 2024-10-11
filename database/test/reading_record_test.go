package test

import (
	"context"
	"testing"
	"time"

	"github.com/aTrejoCoder/reading_tracker_api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRecordInsertAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	readingCollection := getCollection("readings_test")

	defer closeTestDB()

	// Create
	testReading := getTestReading()
	testReading.ReadingStatus = "READING"
	testReading.UpdatedAt = time.Now().Local().UTC()
	testRecord := getTestRecord()
	testReading.ReadingsRecords = []models.ReadingRecord{testRecord}

	// Insert
	_, err := readingCollection.InsertOne(context.TODO(), testReading)
	if err != nil {
		t.Fatalf("Failed to insert reading: %v", err)
	}

	// Get
	var retrievedReading models.Reading
	err = readingCollection.FindOne(context.TODO(), bson.M{"_id": testReading.Id}).Decode(&retrievedReading)
	if err != nil {
		t.Fatalf("Failed to retrieve reading: %v", err)
	}

	// Compare
	assert.NoError(t, err)
	assert.Equal(t, testReading.Id, retrievedReading.Id)
	assert.Equal(t, testReading.ReadingsRecords[0].Id, retrievedReading.ReadingsRecords[0].Id)
	assert.Equal(t, testReading.ReadingsRecords[0].RecordDate.Truncate(time.Millisecond), retrievedReading.ReadingsRecords[0].RecordDate.Truncate(time.Millisecond))
	assert.Equal(t, testReading.ReadingsRecords[0].Notes, retrievedReading.ReadingsRecords[0].Notes)
}

func TestRecordUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	readingCollection := getCollection("readings_test")

	defer closeTestDB()

	// Create initial reading with one record
	testReading := getTestReading()
	testRecord := getTestRecord()
	testReading.ReadingsRecords = []models.ReadingRecord{testRecord}

	// Insert reading
	_, err := readingCollection.InsertOne(context.TODO(), testReading)
	if err != nil {
		t.Fatalf("Failed to insert reading: %v", err)
	}

	// Update the reading record
	filter := bson.M{"_id": testReading.Id, "reading_records._id": testRecord.Id}
	update := bson.M{
		"$set": bson.M{
			"reading_records.$.notes":       "Updated notes",
			"reading_records.$.record_date": time.Now().Local().UTC(),
		},
	}
	_, err = readingCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to update reading record: %v", err)
	}

	// Retrieve the updated reading and check the record
	var updatedReading models.Reading
	err = readingCollection.FindOne(context.TODO(), bson.M{"_id": testReading.Id}).Decode(&updatedReading)
	if err != nil {
		t.Fatalf("Failed to retrieve updated reading: %v", err)
	}

	// Compare the updated fields
	assert.NoError(t, err)
	assert.Equal(t, "Updated notes", updatedReading.ReadingsRecords[0].Notes)
	assert.NotEqual(t, testRecord.RecordDate, updatedReading.ReadingsRecords[0].RecordDate) // Ensure record date was updated
}

func TestRecordDelete(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	readingCollection := getCollection("readings_test")

	defer closeTestDB()

	// Create initial reading with one record
	testReading := getTestReading()
	testRecord := getTestRecord()
	testReading.ReadingsRecords = []models.ReadingRecord{testRecord}

	// Insert reading
	_, err := readingCollection.InsertOne(context.TODO(), testReading)
	if err != nil {
		t.Fatalf("Failed to insert reading: %v", err)
	}

	// Delete the record from reading
	filter := bson.M{"_id": testReading.Id}
	update := bson.M{
		"$pull": bson.M{
			"reading_records": bson.M{
				"_id": testRecord.Id,
			},
		},
	}
	_, err = readingCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to delete reading record: %v", err)
	}

	// Retrieve the reading and ensure the record is deleted
	var updatedReading models.Reading
	err = readingCollection.FindOne(context.TODO(), bson.M{"_id": testReading.Id}).Decode(&updatedReading)
	if err != nil {
		t.Fatalf("Failed to retrieve updated reading: %v", err)
	}

	// Verify that the record was deleted
	assert.NoError(t, err)
	assert.Equal(t, 0, len(updatedReading.ReadingsRecords)) // Ensure the record list is now empty
}

func getTestRecord() models.ReadingRecord {
	return models.ReadingRecord{
		Id:         primitive.NewObjectID(),
		Notes:      "record test notes",
		RecordDate: time.Now().Local().UTC(),
	}
}
