package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

type Timeline struct {
	Id                  entity.DatabaseId `json:"id" bson:"_id"`
	PurchaseId          entity.DatabaseId `json:"purchase_id" bson:"purchase_id"`
	Paid                bool              `json:"paid" bson:"paid"`
	Amount              float64           `json:"amount" bson:"amount"`
	ExpectedPaymentDate time.Time         `json:"expected_payment_date" bson:"expected_payment_date"`
	ActualPaymentDate   time.Time         `json:"actual_payment_date" bson:"actual_payment_date"`
}