package purchase

import "github.com/AyokunlePaul/crud-pay-api/src/pkg/response"

type reader interface {
	Get(*Purchase) *response.BaseResponse
	List(string) ([]Purchase, *response.BaseResponse)
}

type writer interface {
	Create(*Purchase) *response.BaseResponse
	Update(*Purchase) *response.BaseResponse
}

type Repository interface {
	reader
	writer
}

type Manager interface {
	Repository
}
