package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type UseCase interface {
	CreatePurchase(string, *purchase.Purchase) *response.BaseResponse
	UpdatePurchase(string, string, string, float64) *response.BaseResponse
	GetAllPurchasesMadeByUser(string) ([]purchase.Purchase, *response.BaseResponse)
	GetPurchase(string, string) (*purchase.Purchase, *response.BaseResponse)
}
