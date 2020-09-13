package timeline

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

func (manager *manager) Create(timeline *Timeline) *response.BaseResponse {
	return manager.repository.Create(timeline)
}

func (manager *manager) CreateList(timelines []Timeline) *response.BaseResponse {
	return manager.repository.CreateList(timelines)
}

func (manager *manager) Get(timeline *Timeline) *response.BaseResponse {
	return manager.repository.Get(timeline)
}

func (manager *manager) List(purchaseId entity.DatabaseId) ([]Timeline, *response.BaseResponse) {
	return manager.repository.List(purchaseId)
}

func (manager *manager) Update(timeline *Timeline) *response.BaseResponse {
	return manager.repository.Update(timeline)
}
