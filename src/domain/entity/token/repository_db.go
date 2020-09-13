package token

import (
	"context"
	"fmt"
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
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

	if redisSetError := redisClient.Set(redisContext, crudPayToken.AccessUuid, userId, accessTokenExpiration).Err();
		redisSetError != nil {
		logger.Error("error writing access token", redisSetError)
		return response.NewInternalServerError("error creating user")
	}

	if redisSetError := redisClient.Set(redisContext, crudPayToken.RefreshUuid, userId, refreshTokenExpiration).Err();
		redisSetError != nil {
		logger.Error("error writing token", redisSetError)
		return response.NewInternalServerError("error creating user")
	}

	return nil
}

func (repository *repository) Get(accessUuid string) (string, *response.BaseResponse) {
	userId, resultError := redisClient.Get(redisContext, accessUuid).Result()
	if resultError != nil {
		return "", repository.errorService.HandleRedisDbError(resultError)
	}
	return userId, nil
}

func (repository *repository) Update(userId string) (*CrudPayToken, *response.BaseResponse) {
	panic("implement me")
}

func (repository *repository) RefreshToken(refreshToken string) (*CrudPayToken, *response.BaseResponse) {
	panic("implement me")
}
