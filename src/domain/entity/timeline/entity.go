package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

const (
	TypeOneTime     Type = "one-time"
	TypeInstallment Type = "installment"
	TypeRecurring   Type = "recurring"
)

type Type string

type Timeline struct {
	Id                  entity.DatabaseId `json:"id" bson:"_id"`
	PurchaseId          entity.DatabaseId `json:"purchase_id" bson:"purchase_id"`
	Paid                bool              `json:"paid" bson:"paid"`
	Amount              float64           `json:"amount" bson:"amount"`
	ShippingFee         float64           `json:"shipping_fee" bson:"shipping_fee"`
	ExpectedPaymentDate time.Time         `json:"expected_payment_date" bson:"expected_payment_date"`
	ActualPaymentDate   *time.Time        `json:"actual_payment_date,omitempty" bson:"actual_payment_date,omitempty"`
}
