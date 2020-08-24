package mongo_client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

var (
	mongoClient *mongo.Client
)

func init() {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	indexModel := mongo.IndexModel{
		Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	mongoClient, clientError := mongo.Connect(mongoContext, options.Client().ApplyURI("mongodb://mongo:27017"))
	if clientError != nil {
		panic(clientError)
	}

	_, indexError := mongoClient.Database("CrudPay").Collection("users").Indexes().CreateOne(mongoContext, indexModel)
	if indexError != nil {
		panic(indexError)
	}
}

func Get() *mongo.Client {
	return mongoClient
}
