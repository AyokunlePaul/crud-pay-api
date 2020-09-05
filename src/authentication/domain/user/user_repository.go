package user

import "github.com/AyokunlePaul/crud-pay-api/src/utils/response"

type Repository interface {
	CreateUser(User) (*User, *response.BaseResponse)
	Get(User) (*User, *response.BaseResponse)
	Update(User, string) (*User, *response.BaseResponse)
	ResetPassword(string) *response.BaseResponse
	RefreshToken(string) (*User, *response.BaseResponse)
}
