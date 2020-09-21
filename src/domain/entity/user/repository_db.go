package user

import (
	"context"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/infra/database"
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

var collection *mongo.Collection

const from = "timeline"

type repository struct {
	errorService crudPayError.Service
}

func Init() {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userIndexModel := mongo.IndexModel{
		Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}

	collection = database.GetCrudPayDatabase().Collection("users")

	_, indexError := collection.Indexes().CreateOne(mongoContext, userIndexModel)
	if indexError != nil {
		panic(indexError)
	}
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) Create(user *User) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := collection.InsertOne(mongoContext, user)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(from, insertionError)
	}

	return nil
}

func (repository *repository) Get(user *User) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"_id": user.Id},
			{"email": user.Email},
		},
	}
	if getUserError := collection.FindOne(mongoContext, filter).Decode(&user); getUserError != nil {
		return repository.errorService.HandleMongoDbError(from, getUserError)
	}

	return nil
}

func (repository *repository) Update(user *User) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateParameter := bson.D{
		{"$set", bson.D{
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"email", user.Email},
			{"token", user.Token},
			{"refresh_token", user.RefreshToken},
			{"updated_at", user.UpdatedAt},
		}},
	}
	filter := bson.M{"_id": user.Id}
	if _, updateUserError := collection.UpdateOne(mongoContext, filter, updateParameter); updateUserError != nil {
		return repository.errorService.HandleMongoDbError(from, updateUserError)
	}

	return nil
}

func (repository *repository) Delete(userId entity.DatabaseId) *response.BaseResponse {
	panic("implement me")
}

func (repository *repository) Search(query string) (*User, *response.BaseResponse) {
	panic("implement me")
}

func (repository *repository) List() ([]User, *response.BaseResponse) {
	panic("implement me")
}
