package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type UseCase interface {
	CreateProduct(string, *product.Product) *response.BaseResponse
	UpdateProduct(string, string, *product.Product) (*product.Product, *response.BaseResponse)
	SearchProduct(string, string) ([]product.Product, *response.BaseResponse)
	GetProductWithId(string, string) (*product.Product, *response.BaseResponse)
	GetAllProductsCreatedByUserWithId(string) ([]product.Product, *response.BaseResponse)
}
