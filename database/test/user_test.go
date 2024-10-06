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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUserInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := GetTestUser()

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
	assert.Empty(t, retrievedUser.ReadingLists)
	assert.Empty(t, retrievedUser.CustomDocuments)
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

	userTest := GetTestUser()

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
	userTest := GetTestUser()
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

func GetTestUser() models.User {
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
		ReadingLists:    []models.ReadingsList{},
		CustomDocuments: []models.CustomDocument{},
		Roles:           []string{"common_user"},

		LastLogin: now,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func TestCustomDocumentInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	// Insert a test user
	userTest := GetTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Insert CustomDocument linked to the user
	customDoc := GetTestCustomDocument()
	filter := bson.M{"_id": userTest.Id}
	update := bson.M{
		"$push": bson.M{
			"custom_documents": customDoc,
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to insert custom document: %v", err)
	}

	// Retrieve and check the CustomDocument by the user ID
	var retrievedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&retrievedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}

	retrievedDoc := retrievedUser.CustomDocuments[0]
	assert.NoError(t, err)
	assert.Equal(t, customDoc.Title, retrievedDoc.Title)
	assert.Equal(t, customDoc.Author, retrievedDoc.Author)
	assert.Equal(t, customDoc.Description, retrievedDoc.Description)
	assert.Equal(t, customDoc.Content, retrievedDoc.Content)
	assert.Equal(t, customDoc.FileURL, retrievedDoc.FileURL)
	assert.Equal(t, customDoc.URL, retrievedDoc.URL)
	assert.Equal(t, customDoc.Tags, retrievedDoc.Tags)
	assert.Equal(t, customDoc.Category, retrievedDoc.Category)
	assert.Equal(t, customDoc.Version, retrievedDoc.Version)
	assert.Equal(t, customDoc.Status, retrievedDoc.Status)
	assert.Equal(t, customDoc.CreatedAt.Truncate(time.Millisecond), retrievedDoc.CreatedAt.Truncate(time.Millisecond))
	assert.Equal(t, customDoc.UpdatedAt.Truncate(time.Millisecond), retrievedDoc.UpdatedAt.Truncate(time.Millisecond))
}

func TestCustomDocumentUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	// Insert a test user
	userTest := GetTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	// Insert a CustomDocument linked to the user
	customDoc := GetTestCustomDocument()
	filter := bson.M{"_id": userTest.Id}
	update := bson.M{
		"$push": bson.M{
			"custom_documents": customDoc,
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to insert custom document: %v", err)
	}

	// Update the CustomDocument
	newTitle := "Updated Title"
	newVersion := "v2"
	updateFilter := bson.M{
		"_id":                  userTest.Id,
		"custom_documents._id": customDoc.Id,
	}

	updateDoc := bson.M{
		"$set": bson.M{
			"custom_documents.$.title":      newTitle,
			"custom_documents.$.version":    newVersion,
			"custom_documents.$.updated_at": time.Now().UTC(),
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), updateFilter, updateDoc)
	if err != nil {
		t.Fatalf("Failed to update custom document: %v", err)
	}

	// Retrieve and check the updated CustomDocument
	var updatedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&updatedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve updated user: %v", err)
	}

	var updatedDoc models.CustomDocument
	for _, doc := range updatedUser.CustomDocuments {
		if doc.Id == customDoc.Id {
			updatedDoc = doc
			break
		}
	}

	// Assert that the update was successful
	assert.Equal(t, customDoc.Id, updatedDoc.Id)
	assert.Equal(t, newTitle, updatedDoc.Title)
	assert.Equal(t, newVersion, updatedDoc.Version)
	assert.NotEqual(t, customDoc.UpdatedAt.Truncate(time.Millisecond), updatedDoc.UpdatedAt.Truncate(time.Millisecond))
}

func TestCustomDocumentDeletion(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := GetTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	customDoc := GetTestCustomDocument()
	filter := bson.M{"_id": userTest.Id}
	update := bson.M{
		"$push": bson.M{
			"custom_documents": customDoc,
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to insert custom document: %v", err)
	}

	deleteFilter := bson.M{"_id": userTest.Id}
	deleteUpdate := bson.M{
		"$pull": bson.M{
			"custom_documents": bson.M{"_id": customDoc.Id},
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), deleteFilter, deleteUpdate)
	if err != nil {
		t.Fatalf("Failed to delete custom document: %v", err)
	}

	var updatedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&updatedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve updated user: %v", err)
	}

	found := false
	for _, doc := range updatedUser.CustomDocuments {
		if doc.Id == customDoc.Id {
			found = true
			break
		}
	}

	assert.False(t, found, "CustomDocument should be deleted")
}

func GetTestCustomDocument() models.CustomDocument {
	now := time.Now().UTC()
	return models.CustomDocument{
		Id:          primitive.NewObjectID(),
		Title:       "Test Custom Document",
		Author:      "Test Author",
		Description: "This is a test description for a custom document.",
		Content:     "Test content here.",
		FileURL:     "http://example.com/test-file",
		URL:         "http://example.com",
		Tags:        []string{"test", "document"},
		Category:    "test-category",
		CreatedAt:   now,
		UpdatedAt:   now,
		Version:     "v1",
		Status:      "draft",
	}
}

func TestReadingListInsertionAndRetrieval(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := GetTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	readingList := GetTestReadingList()

	filter := bson.M{"_id": userTest.Id}
	update := bson.M{
		"$push": bson.M{
			"reading_lists": readingList,
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to insert reading list: %v", err)
	}

	var retrievedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&retrievedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}
}

func GetTestReadingList() models.ReadingsList {
	now := time.Now().UTC()

	return models.ReadingsList{
		Id:          primitive.NewObjectID(),
		ReadingIds:  []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
		Name:        "Test Reading List",
		Description: "This is a test reading list description.",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func TestReadingListUpdate(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := GetTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	readingList := GetTestReadingList()
	filter := bson.M{"_id": userTest.Id}
	update := bson.M{
		"$push": bson.M{
			"reading_lists": readingList,
		},
	}
	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to insert reading list: %v", err)
	}

	updatedName := "Updated Reading List"
	updatedDescription := "This is an updated test reading list."
	update = bson.M{
		"$set": bson.M{
			"reading_lists.$[list].name":        updatedName,
			"reading_lists.$[list].description": updatedDescription,
			"reading_lists.$[list].updated_at":  time.Now().UTC(),
		},
	}
	arrayFilter := bson.A{
		bson.M{"list._id": readingList.Id},
	}
	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: arrayFilter,
	})

	_, err = userCollection.UpdateOne(context.TODO(), filter, update, updateOptions)
	if err != nil {
		t.Fatalf("Failed to update reading list: %v", err)
	}

	var updatedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&updatedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}

	updatedList := updatedUser.ReadingLists[0]
	assert.Equal(t, updatedName, updatedList.Name)
	assert.Equal(t, updatedDescription, updatedList.Description)
	assert.NotEqual(t, readingList.UpdatedAt.Truncate(time.Millisecond), updatedList.UpdatedAt.Truncate(time.Millisecond))
}

func TestReadingListDeletion(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("Can't connect to database: %v", err)
	}
	userCollection := getCollection("users_test")

	defer closeTestDB()
	defer dropCollection(userCollection)

	userTest := GetTestUser()
	_, err := userCollection.InsertOne(context.TODO(), userTest)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	readingList := GetTestReadingList()
	filter := bson.M{"_id": userTest.Id}
	update := bson.M{
		"$push": bson.M{
			"reading_lists": readingList,
		},
	}
	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to insert reading list: %v", err)
	}

	update = bson.M{
		"$pull": bson.M{
			"reading_lists": bson.M{"_id": readingList.Id},
		},
	}
	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatalf("Failed to delete reading list: %v", err)
	}

	var updatedUser models.User
	err = userCollection.FindOne(context.TODO(), filter).Decode(&updatedUser)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}

	assert.Empty(t, updatedUser.ReadingLists)
}
