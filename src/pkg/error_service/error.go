package error_service

import (
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type crudPayError struct{}

func New() Service {
	return &crudPayError{}
}

func (crudPayError *crudPayError) HandleMongoDbError(err error) *response.BaseResponse {
	logger.Error("mongo error", err)
	if writeException, ok := err.(mongo.WriteException); ok {
		for _, exception := range writeException.WriteErrors {
			switch exception.Code {
			case 11000:
				return response.NewBadRequestError("user with email already exist")
			}
		}
		return response.NewBadRequestError("user already exist")
	}
	switch err {
	case mongo.ErrNoDocuments:
		return response.NewNotFoundError("user doesn't exist")
	case mongo.ErrClientDisconnected:
		return response.NewInternalServerError("internal server error")
	default:
		return response.NewInternalServerError(err.Error())
	}
}

func (crudPayError *crudPayError) HandleRedisDbError(err error) *response.BaseResponse {
	logger.Error("redis error", err)
	switch err {
	case redis.Nil:
		return response.NewUnAuthorizedError()
	default:
		return response.NewInternalServerError(err.Error())
	}
}

func (crudPayError *crudPayError) HandleElasticSearchError(err error) *response.BaseResponse {
	logger.Error("elasticsearch error", err)
	return response.NewInternalServerError(fmt.Sprintf("an error occurred: %s", err.Error()))
}

func (crudPayError *crudPayError) HandlePaystackError(err error) *response.BaseResponse {
	logger.Error("paystack error", err)
	return response.NewInternalServerError(fmt.Sprintf("an error occurred: %s", err.Error()))
}
