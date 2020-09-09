package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/password_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"time"
)

type manager struct {
	repository      Repository
	passwordService password_service.Service
}

func NewManager(repository Repository, passwordService password_service.Service) Manager {
	return &manager{
		repository:      repository,
		passwordService: passwordService,
	}
}

func (manager *manager) Create(user *User) *response.BaseResponse {
	hashedPassword, passwordHashError := manager.passwordService.Generate(user.Password)
	if passwordHashError != nil {
		return response.NewInternalServerError(response.ErrorCreatingUser)
	}
	user.Password = hashedPassword

	return manager.repository.Create(user)
}

func (manager *manager) Update(user *User) *response.BaseResponse {
	user.UpdatedAt = time.Now()
	return manager.repository.Update(user)
}

func (manager *manager) Get(user *User) *response.BaseResponse {
	userPassword := user.Password
	if getUserError := manager.repository.Get(user); getUserError != nil {
		return getUserError
	}
	if passwordComparisonError := manager.passwordService.Compare(user.Password, userPassword); passwordComparisonError != nil {
		return response.NewBadRequestError(response.AuthenticationError)
	}
	return nil
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
