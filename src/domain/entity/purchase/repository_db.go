package purchase

import (
	"context"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/infra/database"
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var collection *mongo.Collection

type repository struct {
	errorService crudPayError.Service
}

func Init() {
	collection = database.GetCrudPayDatabase().Collection("purchases")
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) Create(purchase *Purchase) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := collection.InsertOne(mongoContext, purchase)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(insertionError)
	}
	return nil
}

func (repository *repository) Get(purchase *Purchase) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": purchase.Id,
	}
	if getTransactionError := collection.FindOne(mongoContext, filter).Decode(purchase); getTransactionError != nil {
		return repository.errorService.HandleMongoDbError(getTransactionError)
	}
	return nil
}

func (repository *repository) Update(purchase *Purchase) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateParameter := bson.D{
		{"$set", bson.D{
			{"payment_timeline", purchase.Timeline},
			{"successful", purchase.Successful},
		}},
	}
	filter := bson.M{"_id": purchase.Id}

	if _, updatePurchaseError := collection.UpdateOne(mongoContext, filter, updateParameter); updatePurchaseError != nil {
		return repository.errorService.HandleMongoDbError(updatePurchaseError)
	}
	return nil
}

func (repository *repository) List(userId entity.DatabaseId) ([]Purchase, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"created_by": userId}
	purchaseCursor, purchaseFindError := collection.Find(mongoContext, filter)
	if purchaseFindError != nil {
		return nil, repository.errorService.HandleMongoDbError(purchaseFindError)
	}

	var purchases []Purchase
	if purchasesDecodeError := purchaseCursor.All(mongoContext, &purchases); purchasesDecodeError != nil {
		return nil, repository.errorService.HandleMongoDbError(purchasesDecodeError)
	}

	return purchases, nil
}
