package error_service

import "github.com/AyokunlePaul/crud-pay-api/src/pkg/response"

type Service interface {
	HandleMongoDbError(string, error) *response.BaseResponse
	HandleRedisDbError(error) *response.BaseResponse
	HandleElasticSearchError(error) *response.BaseResponse
	HandlePaystackError(error) *response.BaseResponse
}
