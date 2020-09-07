package utilities

import (
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
)

func HandleElasticSearchError(err error) *response.BaseResponse {
	logger.Error("elasticsearch error", err)
	return response.NewInternalServerError(fmt.Sprintf("an error occurred: %s", err.Error()))
}
