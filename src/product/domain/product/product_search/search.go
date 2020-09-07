package product_search

import (
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
)

type Repository interface {
	Search(string) ([]product.Product, *response.BaseResponse)
}
