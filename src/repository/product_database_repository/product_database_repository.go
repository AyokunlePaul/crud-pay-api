package product_database_repository

import (
	"github.com/AyokunlePaul/crud-pay-api/src/clients/mongo_client"
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
)

var productCollection = mongo_client.Get().Database("CrudPay").Collection("products")

type productRepository struct{}

func New() product.Repository {
	return &productRepository{}
}

func (repository *productRepository) Create(product product.Product) (*product.Product, *response.BaseResponse) {
	panic("implement me")
}

func (repository *productRepository) Get(productId string) (*product.Product, *response.BaseResponse) {
	panic("implement me")
}

func (repository *productRepository) Update(product product.Product) (*product.Product, *response.BaseResponse) {
	panic("implement me")
}

func (repository *productRepository) Search(query string) (*product.Product, *response.BaseResponse) {
	panic("implement me")
}
