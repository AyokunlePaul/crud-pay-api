package user

import "github.com/AyokunlePaul/crud-pay-api/src/utils/response"

type Repository interface {
	CreateUser(User) (*User, *response.BaseResponse)
	Get(User) (*User, *response.BaseResponse)
	Update(User) (*User, *response.BaseResponse)
	ResetPassword(string) *response.BaseResponse
}

type Service interface {
	CreateUser(User) (*User, *response.BaseResponse)
	Get(User) (*User, *response.BaseResponse)
	Update(User) (*User, *response.BaseResponse)
	ResetPassword(string) *response.BaseResponse
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) Service {
	return &service{repository: repository}
}

func (service *service) CreateUser(user User) (*User, *response.BaseResponse) {
	if validationError := user.ValidateUserCreation(); validationError != nil {
		return nil, validationError
	}
	return service.repository.CreateUser(user)
}

func (service *service) Get(user User) (*User, *response.BaseResponse) {
	if validationError := user.ValidateUserLogin(); validationError != nil {
		return nil, validationError
	}
	return service.repository.Get(user)
}

func (service *service) Update(user User) (*User, *response.BaseResponse) {
	panic("implement me")
}

func (service *service) ResetPassword(email string) *response.BaseResponse {
	panic("implement me")
}
