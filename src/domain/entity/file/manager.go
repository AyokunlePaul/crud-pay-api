package file

import (
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

func (manager *manager) Create(userId string, file *CrudPayFile) *response.BaseResponse {
	return manager.repository.Create(userId, file)
}

func (manager *manager) CreateList(userId string, files []*CrudPayFile) *response.BaseResponse {
	return manager.repository.CreateList(userId, files)
}
