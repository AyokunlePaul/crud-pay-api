package timeline

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
	collection = database.GetCrudPayDatabase().Collection("timelines")
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) Create(timeline *Timeline) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := collection.InsertOne(mongoContext, timeline)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(insertionError)
	}

	return nil
}

func (repository *repository) CreateList(timelines []interface{}) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := collection.InsertMany(mongoContext, timelines)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(insertionError)
	}

	return nil
}

func (repository *repository) Get(timeline *Timeline) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": timeline.Id,
	}
	if getTimelineError := collection.FindOne(mongoContext, filter).Decode(timeline); getTimelineError != nil {
		return repository.errorService.HandleMongoDbError(getTimelineError)
	}
	return nil
}

func (repository *repository) Update(timeline *Timeline) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateParameter := bson.D{
		{"$set", bson.D{
			{"actual_payment_date", timeline.ActualPaymentDate},
		}},
	}
	filter := bson.M{"_id": timeline.Id}
	if _, updateTimelineError := collection.UpdateOne(mongoContext, filter, updateParameter); updateTimelineError != nil {
		return repository.errorService.HandleMongoDbError(updateTimelineError)
	}
	return nil
}

func (repository *repository) List(purchaseId entity.DatabaseId) ([]Timeline, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"purchase_id": purchaseId}
	purchaseCursor, purchaseFindError := collection.Find(mongoContext, filter)
	if purchaseFindError != nil {
		return nil, repository.errorService.HandleMongoDbError(purchaseFindError)
	}

	var purchases []Timeline
	if purchasesDecodeError := purchaseCursor.All(mongoContext, &purchases); purchasesDecodeError != nil {
		return nil, repository.errorService.HandleMongoDbError(purchasesDecodeError)
	}

	return purchases, nil
}
