package user

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"time"
)

type reader interface {
	Get(*User) *response.BaseResponse
	Search(string) (*User, *response.BaseResponse)
	List(time.Time, time.Time) (int64, *response.BaseResponse)
	ListAdmin() ([]User, *response.BaseResponse)
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
	IncrementTotalPurchase(entity.DatabaseId) *response.BaseResponse
}
