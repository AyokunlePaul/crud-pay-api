package product_database_repository

import (
	"context"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/token"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/mongo_client"
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var productCollection = mongo_client.Get().Database("CrudPay").Collection("products")

type productRepository struct {
	tokenRepository token.Repository
}

func New(tokenRepository token.Repository) product.Repository {
	return &productRepository{
		tokenRepository: tokenRepository,
	}
}

func (repository *productRepository) Create(product product.Product, token string) (*product.Product, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId, tokenError := repository.tokenRepository.Get(token)
	if tokenError != nil {
		return nil, tokenError
	}

	ownerId, _ := primitive.ObjectIDFromHex(*userId)

	product.Id = primitive.NewObjectID()
	product.OwnerId = ownerId

	_, insertionError := productCollection.InsertOne(mongoContext, product)
	if insertionError != nil {
		return nil, utilities.HandleMongoUserExceptions(insertionError)
	}
	return &product, nil
}

func (repository *productRepository) Get(productId string, token string) (*product.Product, *response.BaseResponse) {
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, tokenError := repository.tokenRepository.Get(token)
	if tokenError != nil {
		return nil, tokenError
	}
	id, _ := primitive.ObjectIDFromHex(productId)
	filter := bson.M{"_id": id}

	var Product *product.Product

	if getProductError := productCollection.FindOne(mongoContext, filter).Decode(&Product); getProductError != nil {
		return nil, utilities.HandleMongoUserExceptions(getProductError)
	}

	return Product, nil
}

func (repository *productRepository) GetProducts(token string) ([]product.Product, *response.BaseResponse) {
	var products []product.Product
	mongoContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId, tokenError := repository.tokenRepository.Get(token)
	if tokenError != nil {
		return nil, tokenError
	}

	ownerId, _ := primitive.ObjectIDFromHex(*userId)
	filter := bson.M{"ownerId": ownerId}

	productsCursor, getProductsError := productCollection.Find(mongoContext, filter)
	if getProductsError != nil {
		return nil, utilities.HandleMongoUserExceptions(getProductsError)
	}

	if productsDecodeError := productsCursor.All(mongoContext, &products); productsDecodeError != nil {
		return nil, utilities.HandleMongoUserExceptions(productsDecodeError)
	}

	return products, nil
}

func (repository *productRepository) Update(product product.Product, token string) (*product.Product, *response.BaseResponse) {
	panic("implement me")
}

func (repository *productRepository) Search(query string, token string) (*product.Product, *response.BaseResponse) {
	panic("implement me")
}
