package authentication

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type UseCase interface {
	Create(*user.User) *response.BaseResponse
	LogIn(*user.User) *response.BaseResponse
	Update(string, user.User) (*user.User, *response.BaseResponse)
	ForgotPassword(string, string) *response.BaseResponse
	RefreshToken(string) (*user.User, *response.BaseResponse)
}
