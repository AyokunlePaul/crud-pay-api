package user

import (
	"context"
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/infra/database"
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

var collection *mongo.Collection

type mongoDbRepository struct {
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
	return &mongoDbRepository{
		errorService: errorService,
	}
}

func (repository *mongoDbRepository) Create(user *User) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := collection.InsertOne(mongoContext, user)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(insertionError)
	}

	return nil
}

func (repository *mongoDbRepository) Get(user *User) *response.BaseResponse {
	logger.Info(fmt.Sprintf("User details: %v", user))
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"_id": user.Id},
			{"email": user.Email},
		},
	}
	if getUserError := collection.FindOne(mongoContext, filter).Decode(&user); getUserError != nil {
		return repository.errorService.HandleMongoDbError(getUserError)
	}

	return nil
}

func (repository *mongoDbRepository) Update(user *User) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateParameter := bson.D{
		{"$set", bson.D{
			{"first_name", user.FirstName},
			{"last_name", user.LastName},
			{"email", user.Email},
			{"updated_at", user.UpdatedAt},
		}},
	}
	filter := bson.M{"_id": user.Id}
	if _, updateUserError := collection.UpdateOne(mongoContext, filter, updateParameter); updateUserError != nil {
		return repository.errorService.HandleMongoDbError(updateUserError)
	}

	return nil
}

func (repository *mongoDbRepository) Delete(userId entity.DatabaseId) *response.BaseResponse {
	panic("implement me")
}

func (repository *mongoDbRepository) Search(query string) (*User, *response.BaseResponse) {
	panic("implement me")
}

func (repository *mongoDbRepository) List() ([]User, *response.BaseResponse) {
	panic("implement me")
}
