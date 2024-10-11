package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

const databaseName = "reading_tracker"

func DbConn() {
	// Connect to docker container
	const uri = "mongodb://reading_tracker-mongo:27017/" + databaseName

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(`"
	////////////////////////////////////////////////////////////
		Successfully connected to MongoDB!
	////////////////////////////////////////////////////////////
	`)

}
