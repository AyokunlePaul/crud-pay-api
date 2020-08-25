package authentication

import (
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleMongoUserExceptions(err error) *response.BaseResponse {
	if writeException, ok := err.(mongo.WriteException); ok {
		for _, exception := range writeException.WriteErrors {
			logger.Error("duplicate exceptions", exception)
			switch exception.Code {
			case 11000:
				return response.NewBadRequestError("user with email already exist")
			}
		}
		return response.NewBadRequestError("duplicate index found")
	}
	switch err {
	case mongo.ErrNoDocuments:
		return response.NewNotFoundError("user doesn't exist")
	default:
		return response.NewInternalServerError(err.Error())
	}
}
