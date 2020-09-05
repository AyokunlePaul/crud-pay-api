package user_service

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
)

type Service interface {
	CreateUser(user.User) (*user.User, *response.BaseResponse)
	Get(user.User) (*user.User, *response.BaseResponse)
	Update(user.User, string) (*user.User, *response.BaseResponse)
	ResetPassword(string) *response.BaseResponse
	RefreshToken(string) (*user.User, *response.BaseResponse)
}

type service struct {
	repository user.Repository
}

func New(repository user.Repository) Service {
	return &service{repository: repository}
}

func (service *service) CreateUser(user user.User) (*user.User, *response.BaseResponse) {
	if validationError := user.ValidateUserCreation(); validationError != nil {
		return nil, validationError
	}
	return service.repository.CreateUser(user)
}

func (service *service) Get(user user.User) (*user.User, *response.BaseResponse) {
	if validationError := user.ValidateUserLogin(); validationError != nil {
		return nil, validationError
	}
	return service.repository.Get(user)
}

func (service *service) Update(user user.User, token string) (*user.User, *response.BaseResponse) {
	return service.repository.Update(user, token)
}

func (service *service) ResetPassword(email string) *response.BaseResponse {
	panic("implement me")
}

func (service *service) RefreshToken(refreshToken string) (*user.User, *response.BaseResponse) {
	return service.repository.RefreshToken(refreshToken)
}
