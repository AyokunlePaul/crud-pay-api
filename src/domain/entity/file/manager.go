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

func (manager *manager) Create(file *CrudPayFile) *response.BaseResponse {
	return manager.repository.Create(file)
}

func (manager *manager) CreateList(files []CrudPayFile) *response.BaseResponse {
	return manager.repository.CreateList(files)
}
