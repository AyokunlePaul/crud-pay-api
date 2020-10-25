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

const from = "user"

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
			{"total_purchase", user.TotalPurchase},
			{"is_deleted", user.IsDeleted},
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
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateParameter := bson.D{
		{"$key", bson.D{
			{"is_deleted", true},
		}},
	}
	filter := bson.M{"_id": userId}
	if _, updateUserError := collection.UpdateOne(mongoContext, filter, updateParameter); updateUserError != nil {
		return repository.errorService.HandleMongoDbError(from, updateUserError)
	}

	return nil
}

func (repository *repository) Search(string) (*User, *response.BaseResponse) {
	panic("implement me")
}

func (repository *repository) ListAdmin() ([]User, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var admins []User

	filter := bson.M{"$and": []bson.M{
		{"is_admin": true},
		{"is_deleted": bson.M{
			"$ne": true,
		}},
	}}
	adminCursor, getAllAdminError := collection.Find(mongoContext, filter)
	if getAllAdminError != nil {
		return nil, repository.errorService.HandleMongoDbError(from, getAllAdminError)
	}

	if adminDecodeError := adminCursor.All(mongoContext, &admins); adminDecodeError != nil {
		return nil, repository.errorService.HandleMongoDbError(from, adminDecodeError)
	}

	return admins, nil
}

func (repository *repository) List(fromDate, toDate time.Time) (int64, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{
		{"created_at", bson.M{
			"$gte": fromDate,
			"$lt":  toDate,
		}},
	}
	if totalCount, countError := collection.CountDocuments(mongoContext, filter); countError != nil {
		return 0, repository.errorService.HandleMongoDbError(from, countError)
	} else {
		return totalCount, nil
	}
}
