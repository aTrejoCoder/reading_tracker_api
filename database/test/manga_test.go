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

func TestMangaInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	mangaCollection := getCollection("manga_test")

	defer closeTestDB()
	defer dropCollection(mangaCollection)

	testManga := GetTestManga()

	// Insert
	_, err := mangaCollection.InsertOne(context.TODO(), testManga)
	if err != nil {
		t.Fatalf("Failed to insert manga: %v", err)
	}

	// Retrieve the manga by ID (_id)
	var retrievedManga models.Manga
	err = mangaCollection.FindOne(context.TODO(), bson.M{"_id": testManga.Id}).Decode(&retrievedManga)
	if err != nil {
		t.Fatalf("Failed to retrieve manga: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, testManga.Id, retrievedManga.Id)
	assert.Equal(t, testManga.Author, retrievedManga.Author)
	assert.Equal(t, testManga.Demography, retrievedManga.Demography) // Check for corrected field
	assert.Equal(t, testManga.CoverImageURL, retrievedManga.CoverImageURL)
	assert.Equal(t, testManga.Volume, retrievedManga.Volume)
	assert.Equal(t, testManga.Chapters, retrievedManga.Chapters)
	assert.Equal(t, testManga.PublicationDate.Truncate(time.Millisecond), retrievedManga.PublicationDate.Truncate(time.Millisecond))
	assert.Equal(t, testManga.Publisher, retrievedManga.Publisher)
	assert.Equal(t, testManga.Description, retrievedManga.Description)
	assert.Equal(t, testManga.Genres, retrievedManga.Genres)
	assert.Equal(t, testManga.CreatedAt.Truncate(time.Millisecond), retrievedManga.CreatedAt.Truncate(time.Millisecond))
	assert.Equal(t, testManga.UpdatedAt.Truncate(time.Millisecond), retrievedManga.UpdatedAt.Truncate(time.Millisecond))
}

func TestMangaUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	mangaCollection := getCollection("manga_test")

	defer closeTestDB()
	defer dropCollection(mangaCollection)

	testManga := GetTestManga()

	// Insert
	_, err := mangaCollection.InsertOne(context.TODO(), testManga)
	if err != nil {
		t.Fatalf("Failed to insert manga: %v", err)
	}

	// Update
	filter := bson.M{"_id": testManga.Id}
	update := bson.M{
		"$set": bson.M{
			"author":     "Updated Author",
			"title":      "Updated Manga Title",
			"chapters":   5,
			"updated_at": time.Now().UTC(),
		},
	}
	_, err = mangaCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to update manga: %v", err)
	}

	// Retrieve the updated manga
	var updatedManga models.Manga
	err = mangaCollection.FindOne(context.TODO(), bson.M{"_id": testManga.Id}).Decode(&updatedManga)
	if err != nil {
		t.Fatalf("Failed to retrieve updated manga: %v", err)
	}

	// Assert that there are no errors
	assert.NoError(t, err)

	// Compare updated fields
	assert.Equal(t, testManga.Id, updatedManga.Id)
	assert.Equal(t, "Updated Author", updatedManga.Author)
	assert.Equal(t, "Updated Manga Title", updatedManga.Title)
	assert.Equal(t, 5, updatedManga.Chapters)

	assert.NotEqual(t, testManga.UpdatedAt.Truncate(time.Millisecond), updatedManga.UpdatedAt.Truncate(time.Millisecond))
}

func TestMangaDeletion(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	mangaCollection := getCollection("manga_test")

	defer closeTestDB()
	defer dropCollection(mangaCollection)

	// Insert
	testManga := GetTestManga()
	_, err := mangaCollection.InsertOne(context.TODO(), testManga)
	if err != nil {
		t.Fatalf("Failed to insert manga: %v", err)
	}

	// Delete
	filter := bson.M{"_id": testManga.Id}
	_, err = mangaCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Fatalf("Failed to delete manga: %v", err)
	}

	// Verify
	var deletedManga models.Manga
	err = mangaCollection.FindOne(context.TODO(), filter).Decode(&deletedManga)

	// Compare
	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}

func GetTestManga() models.Manga {
	now := time.Now().Local().UTC()
	return models.Manga{
		Id:              primitive.NewObjectID(),
		Author:          "Test Author",
		Title:           "Test Manga Title",
		Volume:          1,
		Demography:      "Seinen",
		Chapters:        10,
		CoverImageURL:   "http://example.com/test-cover.jpg",
		PublicationDate: now.AddDate(-2, 0, 0), // 2 years ago
		Publisher:       "Test Publisher",
		Description:     "This is a test description for the manga.",
		Genres:          []string{"Action", "Romance"},
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
