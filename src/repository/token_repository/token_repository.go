package token_repository

import (
	"context"
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/token"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/redis_client"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities"
	"time"
)

var redisContext = context.Background()

type tokenRepository struct{}

func New() token.Repository {
	return &tokenRepository{}
}

func (repository *tokenRepository) CreateToken(userId string) (*token.CrudPayToken, *response.BaseResponse) {
	redisClient := redis_client.Get()

	payToken, createTokenError := utilities.CreateToken(userId)
	if createTokenError != nil {
		return nil, createTokenError
	}

	accessTokenExpiration := time.Unix(payToken.AccessTokenExpires, 0).Sub(time.Now())
	refreshTokenExpiration := time.Unix(payToken.RefreshTokenExpires, 0).Sub(time.Now())

	if redisSetError := redisClient.Set(
		redisContext, payToken.AccessUuid,
		userId, accessTokenExpiration,
	).Err(); redisSetError != nil {
		logger.Error("error writing access token", redisSetError)
		return nil, response.NewInternalServerError("error creating user")
	}
	if redisSetError := redisClient.Set(
		redisContext, payToken.RefreshUuid,
		userId, refreshTokenExpiration,
	).Err(); redisSetError != nil {
		logger.Error("error writing token", redisSetError)
		return nil, response.NewInternalServerError("error creating user")
	}

	return payToken, nil
}

func (repository *tokenRepository) Get(accessToken string) (*string, *response.BaseResponse) {
	redisClient := redis_client.Get()

	accessUuid, tokenMetaError := utilities.GetTokenMetaData(accessToken, true)
	if tokenMetaError != nil {
		return nil, tokenMetaError
	}

	userId, resultError := redisClient.Get(redisContext, *accessUuid).Result()
	if resultError != nil {
		logger.Error("redis error", resultError)
		return nil, response.NewUnAuthorizedError()
	}

	return &userId, nil
}

func (repository *tokenRepository) Update(userId string) (*token.CrudPayToken, *response.BaseResponse) {
	panic("implement me")
}

func (repository *tokenRepository) RefreshToken(refreshToken string) (*token.CrudPayToken, *response.BaseResponse) {
	redisClient := redis_client.Get()

	refreshUuid, tokenMetaError := utilities.GetTokenMetaData(refreshToken, false)
	if tokenMetaError != nil {
		return nil, tokenMetaError
	}

	userId, resultError := redisClient.Get(redisContext, *refreshUuid).Result()
	if resultError != nil {
		logger.Error("get refresh uuid error", resultError)
		return nil, response.NewInternalServerError("error refreshing token")
	}

	deleteError := redisClient.Del(redisContext, *refreshUuid)
	if deleteError != nil {
		logger.Error("get refresh uuid error", resultError)
		return nil, response.NewInternalServerError("error refreshing token")
	}

	return repository.CreateToken(userId)
}
