package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type reader interface {
	Get(entity.DatabaseId) (*Product, *response.BaseResponse)
	List(entity.DatabaseId) ([]Product, *response.BaseResponse)
}

type writer interface {
	Create(*Product) *response.BaseResponse
	Update(*Product) *response.BaseResponse
	Delete(*Product) *response.BaseResponse
}

type Repository interface {
	reader
	writer
}

type Manager interface {
	Repository
}
