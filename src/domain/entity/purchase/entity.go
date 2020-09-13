package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"time"
)

const (
	BiWeekly   Frequency = "bi-weekly"
	Monthly    Frequency = "monthly"
	Quarterly  Frequency = "quarterly"
	BiAnnually Frequency = "bi-annually"
	Annually   Frequency = "annually"
)

const (
	DurationBiWeekly   = 24 * 7 * 2 * time.Hour
	DurationMonthly    = DurationBiWeekly * 2
	DurationQuarterly  = DurationMonthly * 3
	DurationBiAnnually = DurationQuarterly * 2
	DurationAnnually   = DurationBiAnnually * 2
)

type Frequency string

type DeliveryArea string

type Purchase struct {
	Id                   entity.DatabaseId   `json:"id" bson:"_id"`
	ProductId            entity.DatabaseId   `json:"product_id" bson:"product_id"`
	Reference            string              `json:"reference" bson:"reference"`
	Type                 timeline.Type       `json:"payment_type" bson:"payment_type"`
	Email                string              `json:"email" bson:"email"`
	Amount               float64             `json:"amount" bson:"amount"`
	DebitedAmount        float64             `json:"debited_amount" bson:"debited_amount"`
	NumberOfInstallments int                 `json:"number_of_installments,omitempty" bson:"number_of_installments,omitempty"`
	Frequency            Frequency           `json:"payment_frequency" bson:"payment_frequency"`
	DeliveryArea         DeliveryArea        `json:"delivery_area" bson:"delivery_area"`
	Successful           bool                `json:"successful" bson:"successful"`
	CreatedBy            entity.DatabaseId   `json:"created_by" bson:"created_by"`
	AdditionalDetails    string              `json:"additional_details,omitempty" bson:"additional_details,omitempty"`
	Timeline             []timeline.Timeline `json:"timeline" bson:"-"`
	Duration             time.Duration       `json:"-" bson:"-"`
	CreatedAt            time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at" bson:"updated_at"`
}
