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

func TestUserInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := getTestUser()

	// Insert
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Retrieve the user
	var retrievedUser models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"username": userTest.Username}).Decode(&retrievedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, userTest.Username, retrievedUser.Username)
	assert.Equal(t, userTest.Password, retrievedUser.Password)
	assert.Equal(t, userTest.LastLogin.Truncate(time.Millisecond), retrievedUser.LastLogin.Truncate(time.Millisecond))
	assert.Equal(t, userTest.CreatedAt.Truncate(time.Millisecond), retrievedUser.CreatedAt.Truncate(time.Millisecond))
	assert.Equal(t, userTest.UpdatedAt.Truncate(time.Millisecond), retrievedUser.UpdatedAt.Truncate(time.Millisecond))
	assert.Equal(t, userTest.Profile.FullName, retrievedUser.Profile.FullName)
	assert.Empty(t, retrievedUser.ReadingsLists)
	assert.Equal(t, userTest.Profile.Biography, retrievedUser.Profile.Biography)
	assert.Equal(t, userTest.Profile.ProfileImageURL, retrievedUser.Profile.ProfileImageURL)
	assert.Equal(t, userTest.Profile.ProfileCoverURL, retrievedUser.Profile.ProfileCoverURL)
	assert.Equal(t, userTest.Roles[0], retrievedUser.Roles[0])

}

func TestUserUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := getTestUser()

	// Create
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Update
	filter := bson.M{"username": userTest.Username}
	update := bson.M{
		"$set": bson.M{
			"username":          "updatedUsername",
			"profile.biography": "Updated biography",
			"updated_at":        time.Now().UTC(),
		},
	}
	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Retrive
	var updatedUser models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"username": "updatedUsername"}).Decode(&updatedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve updated user: %v", err)
	}

	// Compare
	assert.NoError(t, err)
	assert.Equal(t, "updatedUsername", updatedUser.Username)
	assert.Equal(t, "Updated biography", updatedUser.Profile.Biography)
	assert.NotEqual(t, userTest.UpdatedAt, updatedUser.UpdatedAt)
}

func TestUserDeletion(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	// Insert
	userTest := getTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Delete
	filter := bson.M{"username": userTest.Username}
	_, err = userCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Verify Delete
	var deletedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&deletedUser)

	assert.Error(t, err)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}

func getTestUser() models.User {
	now := time.Now().UTC()
	return models.User{
		Id:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password123",
		Profile: models.Profile{
			FullName:        "Test User",
			Biography:       "This is a test biography.",
			ProfileImageURL: "http://example.com/image.jpg",
			ProfileCoverURL: "http://example.com/cover.jpg",
			CreatedAt:       now,
			UpdatedAt:       now,
		},
		ReadingsLists: []models.ReadingsList{},
		Roles:         []string{"common_user"},

		LastLogin: now,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
