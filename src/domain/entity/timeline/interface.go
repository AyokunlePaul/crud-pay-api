package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type writer interface {
	Create(*Timeline) *response.BaseResponse
	CreateList([]Timeline) *response.BaseResponse
	Update(*Timeline) *response.BaseResponse
}

type reader interface {
	Get(*Timeline) *response.BaseResponse
	List(entity.DatabaseId) ([]Timeline, *response.BaseResponse)
}

type Repository interface {
	writer
	reader
}

type Manager interface {
	Repository
}
