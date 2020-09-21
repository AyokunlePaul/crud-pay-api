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
	accessToken, refreshToken, tokenCreationError := manager.tokenService.Create(
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

func (manager *manager) RefreshToken(crudPayToken *CrudPayToken, refreshToken string, _ string) *response.BaseResponse {
	refreshUuid, getMetaDataError := manager.tokenService.GetTokenMetaData(refreshToken, false)
	if getMetaDataError != nil {
		return getMetaDataError
	}

	userId, getUserIdError := manager.repository.Get(refreshUuid)
	if getUserIdError != nil {
		return getUserIdError
	}

	newAccessToken, newRefreshToken, tokenCreationError := manager.tokenService.Create(
		crudPayToken.AccessTokenExpires, crudPayToken.RefreshTokenExpires,
		crudPayToken.AccessUuid, crudPayToken.RefreshUuid, userId,
	)
	if tokenCreationError != nil {
		return tokenCreationError
	}

	crudPayToken.AccessToken = newAccessToken
	crudPayToken.RefreshToken = newRefreshToken

	return manager.repository.RefreshToken(crudPayToken, userId, refreshUuid)
}
