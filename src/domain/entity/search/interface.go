package search

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type Repository interface {
	search(Param) (interface{}, *response.BaseResponse)
}

type Manager interface {
	Repository
	SearchProduct(Param) ([]product.Product, *response.BaseResponse)
	SearchUser(Param) ([]user.User, *response.BaseResponse)
}
