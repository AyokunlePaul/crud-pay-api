package file

import "github.com/AyokunlePaul/crud-pay-api/src/pkg/response"

type writer interface {
	Create(*CrudPayFile) *response.BaseResponse
	CreateList([]CrudPayFile) *response.BaseResponse
}

type Repository interface {
	writer
}

type Manager interface {
	Repository
}
