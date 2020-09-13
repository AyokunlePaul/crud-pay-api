package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"time"
)

const (
	TypeOneTime     Type = "one-time"
	TypeInstallment Type = "installment"
	TypeRecurring   Type = "recurring"
)

const (
	BiWeekly   PaymentFrequency = "bi-weekly"
	Monthly    PaymentFrequency = "monthly"
	Quarterly  PaymentFrequency = "quarterly"
	BiAnnually PaymentFrequency = "bi-annually"
	Annually   PaymentFrequency = "annually"
)

const (
	DurationBiWeekly   = 24 * 7 * 2 * time.Hour
	DurationMonthly    = DurationBiWeekly * 2
	DurationQuarterly  = DurationMonthly * 3
	DurationBiAnnually = DurationQuarterly * 2
	DurationAnnually   = DurationBiAnnually * 2
)

type PaymentFrequency string

type Type string

type Purchase struct {
	Id                   entity.DatabaseId   `json:"id" bson:"_id"`
	ProductId            entity.DatabaseId   `json:"product_id" bson:"product_id"`
	Reference            string              `json:"reference" bson:"reference"`
	Type                 Type                `json:"payment_type" bson:"payment_type"`
	Email                string              `json:"email" bson:"email"`
	Amount               float64             `json:"amount" bson:"amount"`
	NumberOfInstallments int64               `json:"number_of_installments" bson:"number_of_installments"`
	Frequency            PaymentFrequency    `json:"payment_frequency" bson:"payment_frequency"`
	Successful           bool                `json:"successful" bson:"successful"`
	CreatedBy            entity.DatabaseId   `json:"created_by" bson:"created_by"`
	AdditionalDetails    string              `json:"additional_details,omitempty" bson:"additional_details,omitempty"`
	PaymentTimelines     []timeline.Timeline `json:"payment_timelines" bson:"payment_timelines"`
	Duration             time.Duration       `json:"-" bson:"-"`
	CreatedAt            time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at" bson:"updated_at"`
}
