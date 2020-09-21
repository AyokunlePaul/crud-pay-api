package token

import (
	"context"
	"fmt"
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

var (
	redisContext = context.Background()
	redisClient  *redis.Client
	redisName    = "REDIS_CONTAINER_NAME"
)

type repository struct {
	errorService crudPayError.Service
}

func init() {
	redisName := os.Getenv(redisName)

	redisClient = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%s", redisName, "6379"),
		DB:         0,
		MaxRetries: 5,
	})
	_, pingResult := redisClient.Ping(context.Background()).Result()
	if pingResult != nil {
		panic(pingResult)
	}
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) CreateToken(crudPayToken *CrudPayToken, userId string) *response.BaseResponse {
	accessTokenExpiration := time.Unix(crudPayToken.AccessTokenExpires, 0).Sub(time.Now())
	refreshTokenExpiration := time.Unix(crudPayToken.RefreshTokenExpires, 0).Sub(time.Now())

	if redisSetError := redisClient.Set(redisContext, crudPayToken.AccessUuid, userId, accessTokenExpiration).Err(); redisSetError != nil {
		return repository.errorService.HandleRedisDbError(redisSetError)
	}

	if redisSetError := redisClient.Set(redisContext, crudPayToken.RefreshUuid, userId, refreshTokenExpiration).Err(); redisSetError != nil {
		return repository.errorService.HandleRedisDbError(redisSetError)
	}

	return nil
}

func (repository *repository) Get(tokenUuid string) (string, *response.BaseResponse) {
	userId, resultError := redisClient.Get(redisContext, tokenUuid).Result()
	if resultError != nil {
		return "", repository.errorService.HandleRedisDbError(resultError)
	}
	return userId, nil
}

func (repository *repository) RefreshToken(crudPayToken *CrudPayToken, userId string, oldTokenUuid string) *response.BaseResponse {
	accessTokenExpiration := time.Unix(crudPayToken.AccessTokenExpires, 0).Sub(time.Now())
	refreshTokenExpiration := time.Unix(crudPayToken.RefreshTokenExpires, 0).Sub(time.Now())

	if redisSetError := redisClient.Set(redisContext, crudPayToken.AccessUuid, userId, accessTokenExpiration).Err(); redisSetError != nil {
		return repository.errorService.HandleRedisDbError(redisSetError)
	}
	if redisDeleteError := redisClient.Del(redisContext, oldTokenUuid).Err(); redisDeleteError != nil {
		return repository.errorService.HandleRedisDbError(redisDeleteError)
	}
	if redisSetError := redisClient.Set(redisContext, crudPayToken.RefreshUuid, userId, refreshTokenExpiration).Err(); redisSetError != nil {
		return repository.errorService.HandleRedisDbError(redisSetError)
	}
	return nil
}
