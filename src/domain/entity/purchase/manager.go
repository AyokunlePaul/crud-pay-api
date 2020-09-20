package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type manager struct {
	repository Repository
}

func NewManager(repository Repository) Manager {
	return &manager{
		repository: repository,
	}
}

func (manager *manager) Create(purchase *Purchase) *response.BaseResponse {
	return manager.repository.Create(purchase)
}

func (manager *manager) Get(purchase *Purchase) *response.BaseResponse {
	return manager.repository.Get(purchase)
}

func (manager *manager) Update(purchase *Purchase) *response.BaseResponse {
	return manager.repository.Update(purchase)
}

func (manager *manager) UpdateTimeline(purchase *Purchase) *response.BaseResponse {
	return manager.repository.UpdateTimeline(purchase)
}

func (manager *manager) List(userId entity.DatabaseId) ([]Purchase, *response.BaseResponse) {
	return manager.repository.List(userId)
}
