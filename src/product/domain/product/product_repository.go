package product

import "github.com/AyokunlePaul/crud-pay-api/src/utils/response"

type Repository interface {
	Create(Product, string) (*Product, *response.BaseResponse)
	Get(string, string) (*Product, *response.BaseResponse)
	GetProducts(string) ([]Product, *response.BaseResponse)
	Update(Product, string) (*Product, *response.BaseResponse)
	Search(string, string) ([]Product, *response.BaseResponse)
}
