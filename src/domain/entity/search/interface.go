package search

import (
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type Repository interface {
	Search(Param) (interface{}, *response.BaseResponse)
}

type Manager interface {
	Repository
}
