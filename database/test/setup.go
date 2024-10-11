package test

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testDatabaseName = "reading_tracker_test"
const testMongoURI = "mongodb://localhost:27019/" + testDatabaseName

var client *mongo.Client // Hacer el cliente accesible globalmente

func TestDbConn(t *testing.T) {
	if err := connectToTestDB(); err != nil {
		t.Fatalf("An error occurred connecting to database: %v", err)
	}
	t.Log("Successfully connected to database")

	if err := closeTestDB(); err != nil {
		t.Fatalf("An error occurred disconnecting from test database: %v", err)
	}
}

func connectToTestDB() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(testMongoURI).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func closeTestDB() error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}

func getCollection(collectionName string) *mongo.Collection {
	return client.Database(testDatabaseName).Collection(collectionName)
}

func dropCollection(collection *mongo.Collection) {
	collection.Drop(context.TODO())
}
