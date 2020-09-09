package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var (
	database *mongo.Database
	mongoUri = os.Getenv("MONGO_CLIENT_URI")
)

func Init() {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, clientError := mongo.Connect(mongoContext, options.Client().ApplyURI(mongoUri))
	if clientError != nil {
		panic(clientError)
	}

	database = mongoClient.Database("CrudPay")
}

func GetCrudPayDatabase() *mongo.Database {
	return database
}
