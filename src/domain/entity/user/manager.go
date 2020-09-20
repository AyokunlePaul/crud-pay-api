package user

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

func (manager *manager) Create(user *User) *response.BaseResponse {
	return manager.repository.Create(user)
}

func (manager *manager) Update(user *User) *response.BaseResponse {
	user.UpdatedAt = time.Now()
	return manager.repository.Update(user)
}

func (manager *manager) Get(user *User) *response.BaseResponse {
	return manager.repository.Get(user)
}

func (manager *manager) Delete(userId entity.DatabaseId) *response.BaseResponse {
	panic("implement me")
}

func (manager *manager) Search(query string) (*User, *response.BaseResponse) {
	panic("implement me")
}

func (manager *manager) List() ([]User, *response.BaseResponse) {
	panic("implement me")
}
