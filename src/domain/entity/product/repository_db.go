package product

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

var (
	productCollection *mongo.Collection
)

type mongoDbRepository struct {
	errorService crudPayError.Service
}

func Init() {
	productCollection = database.GetCrudPayDatabase().Collection("products")
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &mongoDbRepository{
		errorService: errorService,
	}
}

func (repository *mongoDbRepository) Create(product *Product) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := productCollection.InsertOne(mongoContext, product)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(insertionError)
	}
	return nil
}

func (repository *mongoDbRepository) Get(productId entity.DatabaseId) (*Product, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": productId}
	var Product *Product

	if getProductError := productCollection.FindOne(mongoContext, filter).Decode(&Product); getProductError != nil {
		return nil, repository.errorService.HandleMongoDbError(getProductError)
	}

	return Product, nil
}

func (repository *mongoDbRepository) List(ownerId entity.DatabaseId) ([]Product, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []Product
	filter := bson.M{"ownerId": ownerId}

	productsCursor, getProductsError := productCollection.Find(mongoContext, filter)
	if getProductsError != nil {
		return nil, repository.errorService.HandleMongoDbError(getProductsError)
	}

	if productsDecodeError := productsCursor.All(mongoContext, &products); productsDecodeError != nil {
		return nil, repository.errorService.HandleMongoDbError(productsDecodeError)
	}

	return products, nil
}

func (repository *mongoDbRepository) Update(product *Product) *response.BaseResponse {
	panic("implement me")
}

func (repository *mongoDbRepository) Delete(token string, productId string) {
	panic("implement me")
}
