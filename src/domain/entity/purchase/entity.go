package purchase

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

type Purchase struct {
	Id                entity.DatabaseId `json:"id" bson:"_id"`
	ProductId         entity.DatabaseId `json:"product_id" bson:"product_id"`
	ProductName       string            `json:"product_name" bson:"product_name"`
	Reference         string            `json:"reference" bson:"reference"`
	Type              Type              `json:"payment_type" bson:"payment_type"`
	Email             string            `json:"email" bson:"email"`
	Amount            float64           `json:"amount" bson:"amount"`
	Successful        bool              `json:"successful" bson:"successful"`
	CreatedBy         entity.DatabaseId `json:"created_by" bson:"created_by"`
	AdditionalDetails string            `json:"additional_details,omitempty" bson:"additional_details,omitempty"`
	PaymentTimelines  []Timeline        `json:"payment_timelines" bson:"payment_timelines"`
	CreatedAt         time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at" bson:"updated_at"`
}

type Timeline struct {
	Id                  entity.DatabaseId `json:"timeline_id" bson:"timeline_id"`
	Paid                bool              `json:"paid" bson:"paid"`
	Amount              float64           `json:"amount" bson:"amount"`
	ExpectedPaymentDate time.Time         `json:"expected_payment_date" bson:"expected_payment_date"`
	ActualPaymentDate   time.Time         `json:"actual_payment_date" bson:"actual_payment_date"`
}
