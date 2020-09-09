package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"time"
)

type manager struct {
	repository Repository
}

func NewManager(repository Repository) Manager {
	return &manager{
		repository: repository,
	}
}

func (manager *manager) Create(product *Product) *response.BaseResponse {
	return manager.repository.Create(product)
}

func (manager *manager) Get(productId entity.DatabaseId) (*Product, *response.BaseResponse) {
	return manager.repository.Get(productId)
}

func (manager *manager) GetProducts(ownerId entity.DatabaseId) ([]Product, *response.BaseResponse) {
	return manager.repository.GetProducts(ownerId)
}

func (manager *manager) Update(product *Product) *response.BaseResponse {
	product.UpdatedAt = time.Now()
	return manager.repository.Update(product)
}

func (manager *manager) Delete(token string, productId string) {
	panic("implement me")
}
