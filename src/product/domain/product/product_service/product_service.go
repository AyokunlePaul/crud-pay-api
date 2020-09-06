package product_service

import (
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
)

type service struct {
	repository product.Repository
}

type Service interface {
	Create(product.Product, string) (*product.Product, *response.BaseResponse)
	Get(string, string) (*product.Product, *response.BaseResponse)
	GetProducts(string) ([]product.Product, *response.BaseResponse)
	Update(product.Product, string) (*product.Product, *response.BaseResponse)
	Search(string, string) (*product.Product, *response.BaseResponse)
}

func New(repository product.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service *service) Create(product product.Product, token string) (*product.Product, *response.BaseResponse) {
	if validationError := product.IsValidProduct(); validationError != nil {
		return nil, validationError
	}
	return service.repository.Create(product, token)
}

func (service *service) Get(productId string, token string) (*product.Product, *response.BaseResponse) {
	return service.repository.Get(productId, token)
}

func (service *service) GetProducts(token string) ([]product.Product, *response.BaseResponse) {
	return service.repository.GetProducts(token)
}

func (service *service) Update(product product.Product, token string) (*product.Product, *response.BaseResponse) {
	return service.repository.Update(product, token)
}

func (service *service) Search(query string, token string) (*product.Product, *response.BaseResponse) {
	return service.repository.Search(query, token)
}
