package product

import "github.com/AyokunlePaul/crud-pay-api/src/utils/response"

type Repository interface {
	Create(Product) (*Product, *response.BaseResponse)
	Get(string) (*Product, *response.BaseResponse)
	Update(Product) (*Product, *response.BaseResponse)
	Search(string) (*Product, *response.BaseResponse)
}
