package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type reader interface {
	Get(*Purchase) *response.BaseResponse
	List(entity.DatabaseId) ([]Purchase, *response.BaseResponse)
}

type writer interface {
	Create(*Purchase) *response.BaseResponse
	Update(*Purchase) *response.BaseResponse
	UpdateTimeline(*Purchase) *response.BaseResponse
}

type Repository interface {
	reader
	writer
}

type Manager interface {
	Repository
}
