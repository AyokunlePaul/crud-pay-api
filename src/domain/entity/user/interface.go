package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type reader interface {
	Get(*User) *response.BaseResponse
	Search(string) (*User, *response.BaseResponse)
	List() ([]User, *response.BaseResponse)
}

type writer interface {
	Create(*User) *response.BaseResponse
	Update(*User) *response.BaseResponse
	Delete(entity.DatabaseId) *response.BaseResponse
}

type Repository interface {
	writer
	reader
}

type Manager interface {
	Repository
}
