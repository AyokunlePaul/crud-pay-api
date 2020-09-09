package token

import "github.com/AyokunlePaul/crud-pay-api/src/pkg/response"

type reader interface {
	Get(string) (string, *response.BaseResponse)
}

type writer interface {
	CreateToken(*CrudPayToken, string) *response.BaseResponse
	Update(string) (*CrudPayToken, *response.BaseResponse)
	RefreshToken(string) (*CrudPayToken, *response.BaseResponse)
}

type Repository interface {
	reader
	writer
}

type Manager interface {
	Repository
}
