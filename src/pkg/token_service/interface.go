package token_service

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type Service interface {
	Create(int64, int64, string, string, string) (string, string, *response.BaseResponse)
	VerifyAndExtract(string, bool) (*entity.CrudPayJwtToken, *response.BaseResponse)
	CheckTokenValidity(string, bool) *response.BaseResponse
	GetTokenMetaData(string, bool) (string, *response.BaseResponse)
}
