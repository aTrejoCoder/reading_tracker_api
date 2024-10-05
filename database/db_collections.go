package database

import "go.mongodb.org/mongo-driver/mongo"

const readingCollection = "readings"
const userCollection = "users"
const booksCollection = "books"
const mangaCollection = "mangas"

func GetReadingColletion() *mongo.Collection {
	return Client.Database(databaseName).Collection(readingCollection)
}

func GetUserColletion() *mongo.Collection {
	return Client.Database(databaseName).Collection(userCollection)
}

func GetBookColletion() *mongo.Collection {
	return Client.Database(databaseName).Collection(booksCollection)
}

func GetMangaColletion() *mongo.Collection {
	return Client.Database(databaseName).Collection(mangaCollection)
}
