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

const from = "purchase"

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
		return repository.errorService.HandleMongoDbError(from, insertionError)
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
		return repository.errorService.HandleMongoDbError(from, getTransactionError)
	}
	return nil
}

func (repository *repository) Update(purchase *Purchase) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateParameter := bson.D{
		{"$set", bson.D{
			{"timeline", purchase.Timeline},
			{"reference", purchase.Reference},
			{"successful", purchase.Successful},
		}},
	}
	filter := bson.M{"_id": purchase.Id}

	if _, updatePurchaseError := collection.UpdateOne(mongoContext, filter, updateParameter); updatePurchaseError != nil {
		return repository.errorService.HandleMongoDbError(from, updatePurchaseError)
	}
	return nil
}

func (repository *repository) List(userId entity.DatabaseId) ([]Purchase, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"created_by": userId}
	purchaseCursor, purchaseFindError := collection.Find(mongoContext, filter)
	if purchaseFindError != nil {
		return nil, repository.errorService.HandleMongoDbError(from, purchaseFindError)
	}

	var purchases []Purchase
	if purchasesDecodeError := purchaseCursor.All(mongoContext, &purchases); purchasesDecodeError != nil {
		return nil, repository.errorService.HandleMongoDbError(from, purchasesDecodeError)
	}

	return purchases, nil
}

func (repository *repository) UpdateTimeline(purchase *Purchase) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": purchase.Id}
	updateQuery := bson.D{{"$set", bson.D{
		{"successful", purchase.Successful},
		{"updated_at", time.Now()},
		{"timeline.$.paid", true},
		{"timeline.$.actual_payment_date", time.Now()},
	}}}
	if _, updatePurchaseError := collection.UpdateOne(mongoContext, filter, updateQuery); updatePurchaseError != nil {
		return repository.errorService.HandleMongoDbError(from, updatePurchaseError)
	}
	return repository.Get(purchase)
}

func (repository *repository) ListData(fromDate, toDate time.Time) (int64, *response.BaseResponse) {
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
