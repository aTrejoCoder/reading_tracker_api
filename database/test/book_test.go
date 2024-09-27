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

func TestBookInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	bookCollection := getCollection("books_test")

	defer closeTestDB()
	defer dropCollection(bookCollection)

	testBook := getTestBook()

	// Insert
	_, err := bookCollection.InsertOne(context.TODO(), testBook)
	if err != nil {
		t.Fatalf("Failed to insert books: %v", err)
	}

	// Retrieve the books by ID (_id)
	var retrievedBook models.Book
	err = bookCollection.FindOne(context.TODO(), bson.M{"_id": testBook.Id}).Decode(&retrievedBook)
	if err != nil {
		t.Fatalf("Failed to retrieve reading: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, testBook.Id, retrievedBook.Id)
	assert.Equal(t, testBook.Author, retrievedBook.Author)
	assert.Equal(t, testBook.ISBN, retrievedBook.ISBN)
	assert.Equal(t, testBook.Name, retrievedBook.Name)
	assert.Equal(t, testBook.CoverImageURL, retrievedBook.CoverImageURL)
	assert.Equal(t, testBook.Edition, retrievedBook.Edition)
	assert.Equal(t, testBook.Pages, retrievedBook.Pages)
	assert.Equal(t, testBook.Language, retrievedBook.Language)
	assert.Equal(t, testBook.PublicationDate.Truncate(time.Millisecond), retrievedBook.PublicationDate.Truncate(time.Millisecond))
	assert.Equal(t, testBook.Publisher, retrievedBook.Publisher)
	assert.Equal(t, testBook.Description, retrievedBook.Description)
	assert.Equal(t, testBook.Genres, retrievedBook.Genres)
	assert.Equal(t, testBook.CreatedAt.Truncate(time.Millisecond), retrievedBook.CreatedAt.Truncate(time.Millisecond))
	assert.Equal(t, testBook.UpdatedAt.Truncate(time.Millisecond), retrievedBook.UpdatedAt.Truncate(time.Millisecond))
}

func TestBookUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	bookCollection := getCollection("books_test")

	defer closeTestDB()
	defer dropCollection(bookCollection)

	testBook := getTestBook()

	// Insert
	_, err := bookCollection.InsertOne(context.TODO(), testBook)
	if err != nil {
		t.Fatalf("Failed to insert book: %v", err)
	}

	// Update
	filter := bson.M{"_id": testBook.Id}
	update := bson.M{
		"$set": bson.M{
			"author":     "Updated Author",
			"name":       "Updated Book Title",
			"pages":      500,
			"updated_at": time.Now().UTC(),
		},
	}
	_, err = bookCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to update book: %v", err)
	}

	// Retrieve the updated book
	var updatedBook models.Book
	err = bookCollection.FindOne(context.TODO(), bson.M{"_id": testBook.Id}).Decode(&updatedBook)
	if err != nil {
		t.Fatalf("Failed to retrieve updated book: %v", err)
	}

	// Assert that there are no errors
	assert.NoError(t, err)

	// Compare updated fields
	assert.Equal(t, testBook.Id, updatedBook.Id)
	assert.Equal(t, "Updated Author", updatedBook.Author)
	assert.Equal(t, "Updated Book Title", updatedBook.Name)
	assert.Equal(t, 500, updatedBook.Pages)

	assert.NotEqual(t, testBook.UpdatedAt.Truncate(time.Millisecond), updatedBook.UpdatedAt.Truncate(time.Millisecond))
}

func TestBookDeletion(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	bookCollection := getCollection("books_test")

	defer closeTestDB()
	defer dropCollection(bookCollection)

	// Insert
	testBook := getTestBook()
	_, err := bookCollection.InsertOne(context.TODO(), testBook)
	if err != nil {
		t.Fatalf("Failed to insert book: %v", err)
	}

	// Delete
	filter := bson.M{"_id": testBook.Id}
	_, err = bookCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Fatalf("Failed to delete book: %v", err)
	}

	// Verify
	var deletedBook models.Book
	err = bookCollection.FindOne(context.TODO(), filter).Decode(&deletedBook)

	// Compare
	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}

func getTestBook() models.Book {
	now := time.Now().Local().UTC()
	return models.Book{
		Id:              primitive.NewObjectID(),
		Author:          "Test Author",
		ISBN:            "123-4567890123",
		Name:            "Test Book Title",
		CoverImageURL:   "http://example.com/test-cover.jpg",
		Edition:         "1st Edition",
		Pages:           350,
		Language:        "English",
		PublicationDate: now.AddDate(-2, 0, 0), // 2 year ago
		Publisher:       "Test Publisher",
		Description:     "This is a test description for the book.",
		Genres:          []string{"Fiction", "Adventure"},
		ReadingList:     []primitive.ObjectID{},
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}
