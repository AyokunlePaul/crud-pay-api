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
	clientUri   = "mongodb+srv://admin:admin@crudpay-api-cluster.0zbrq.gcp.mongodb.net/<dbname>?retryWrites=true&w=majority"
)

func init() {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	indexModel := mongo.IndexModel{
		Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	mongoClient, clientError := mongo.Connect(mongoContext, options.Client().ApplyURI(clientUri))
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
