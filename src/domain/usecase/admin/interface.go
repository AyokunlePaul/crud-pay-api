package admin

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type UseCase interface {
	CreateNew(*user.User) *response.BaseResponse
	GetAll() ([]user.User, *response.BaseResponse)
	GetDailyStat() (map[string]interface{}, *response.BaseResponse)
	Delete(*user.User) *response.BaseResponse
	Search(string) ([]user.User, *response.BaseResponse)
}
