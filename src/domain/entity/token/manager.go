package token

import (
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/token_service"
)

type manager struct {
	repository   Repository
	tokenService token_service.Service
}

func NewManager(repository Repository, tokenService token_service.Service) Manager {
	return &manager{
		repository:   repository,
		tokenService: tokenService,
	}
}

func (manager *manager) Get(accessToken string) (string, *response.BaseResponse) {
	accessUuid, tokenMetaError := manager.tokenService.GetTokenMetaData(accessToken, true)
	if tokenMetaError != nil {
		return "", tokenMetaError
	}

	return manager.repository.Get(accessUuid)
}

func (manager *manager) CreateToken(crudPayToken *CrudPayToken, userId string) *response.BaseResponse {
	accessToken, refreshToken, tokenCreationError :=
		manager.tokenService.Create(
			crudPayToken.AccessTokenExpires, crudPayToken.RefreshTokenExpires,
			crudPayToken.AccessUuid, crudPayToken.RefreshUuid, userId,
		)
	if tokenCreationError != nil {
		return tokenCreationError
	}

	crudPayToken.AccessToken = accessToken
	crudPayToken.RefreshToken = refreshToken
	return manager.repository.CreateToken(crudPayToken, userId)
}

func (manager *manager) Update(userId string) (*CrudPayToken, *response.BaseResponse) {
	panic("implement me")
}

func (manager *manager) RefreshToken(refreshToken string) (*CrudPayToken, *response.BaseResponse) {
	panic("implement me")
}
