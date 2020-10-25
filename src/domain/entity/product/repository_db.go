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

var collection *mongo.Collection

const from = "product"

type repository struct {
	errorService crudPayError.Service
}

func Init() {
	collection = database.GetCrudPayDatabase().Collection("products")
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) Create(product *Product) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, insertionError := collection.InsertOne(mongoContext, product)
	if insertionError != nil {
		return repository.errorService.HandleMongoDbError(from, insertionError)
	}
	return nil
}

func (repository *repository) Get(productId entity.DatabaseId) (*Product, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": productId}
	var Product *Product

	if getProductError := collection.FindOne(mongoContext, filter).Decode(&Product); getProductError != nil {
		return nil, repository.errorService.HandleMongoDbError(from, getProductError)
	}

	return Product, nil
}

func (repository *repository) List(ownerId entity.DatabaseId) ([]Product, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var products []Product
	filter := bson.M{"owner_id": ownerId}

	productsCursor, getProductsError := collection.Find(mongoContext, filter)
	if getProductsError != nil {
		return nil, repository.errorService.HandleMongoDbError("product", getProductsError)
	}

	if productsDecodeError := productsCursor.All(mongoContext, &products); productsDecodeError != nil {
		return nil, repository.errorService.HandleMongoDbError("product", productsDecodeError)
	}

	return products, nil
}

func (repository *repository) Update(product *Product) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": product.Id}
	updateParameter := bson.D{
		{"$set", bson.D{
			{"product_name", product.Name},
			{"allow_installment", product.AllowInstallment},
			{"delivery_areas", product.DeliveryAreas},
			{"delivery_groups", product.DeliveryGroups},
			{"max_installment", product.MaxInstallment},
			{"amount", product.Amount},
			{"payment_frequencies", product.PaymentFrequencies},
			{"pictures", product.Pictures},
			{"updated_at", product.UpdatedAt},
		}},
	}
	if _, productUpdateError := collection.UpdateOne(mongoContext, filter, updateParameter); productUpdateError != nil {
		return repository.errorService.HandleMongoDbError(from, productUpdateError)
	}
	return nil
}

func (repository *repository) Delete(product *Product) *response.BaseResponse {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": product.Id}
	updateParameter := bson.D{
		{"$set", bson.D{
			{"is_deleted", product.IsDeleted},
			{"updated_at", product.UpdatedAt},
		}},
	}
	if _, productUpdateError := collection.UpdateOne(mongoContext, filter, updateParameter); productUpdateError != nil {
		return repository.errorService.HandleMongoDbError(from, productUpdateError)
	}
	return nil
}
