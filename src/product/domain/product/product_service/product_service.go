package product_service

import (
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
)

type service struct {
	repository product.Repository
}

type Service interface {
	Create(product.Product) (*product.Product, *response.BaseResponse)
	Get(string) (*product.Product, *response.BaseResponse)
	Update(product.Product) (*product.Product, *response.BaseResponse)
	Search(string) (*product.Product, *response.BaseResponse)
}

func New(repository product.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service *service) Create(product product.Product) (*product.Product, *response.BaseResponse) {

	return service.repository.Create(product)
}

func (service *service) Get(productId string) (*product.Product, *response.BaseResponse) {
	return service.repository.Get(productId)
}

func (service *service) Update(product product.Product) (*product.Product, *response.BaseResponse) {
	return service.repository.Update(product)
}

func (service *service) Search(query string) (*product.Product, *response.BaseResponse) {
	return service.repository.Search(query)
}
