package token

import "github.com/AyokunlePaul/crud-pay-api/src/utils/response"

type Repository interface {
	CreateToken(string) (*CrudPayToken, *response.BaseResponse)
	Get(string) (*string, *response.BaseResponse)
	Update(string) (*CrudPayToken, *response.BaseResponse)
	RefreshToken(string) (*CrudPayToken, *response.BaseResponse)
}
