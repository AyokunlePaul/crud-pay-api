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

func NewTokenRepository() token.Repository {
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
	}
	if redisSetError := redisClient.Set(
		redisContext, payToken.RefreshUuid,
		userId, refreshTokenExpiration,
	).Err(); redisSetError != nil {
		logger.Error("error writing token", redisSetError)
	}

	return payToken, nil
}

func (repository *tokenRepository) Get(accessToken string) (*string, *response.BaseResponse) {
	accessUuid, tokenMetaError := utilities.GetTokenMetaData(accessToken)
	if tokenMetaError != nil {
		return nil, tokenMetaError
	}

	redisClient := redis_client.Get()
	userId, resultError := redisClient.Get(redisContext, *accessUuid).Result()
	if resultError != nil {
		return nil, response.NewUnAuthorizedError()
	}

	return &userId, nil
}

func (repository *tokenRepository) Update(userId string) (*token.CrudPayToken, *response.BaseResponse) {
	panic("implement me")
}
