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

type Group struct {
	Name        string   `json:"name" bson:"name"`
	Areas       []string `json:"delivery_areas" bson:"areas"`
	ShippingFee float64  `json:"shipping_fee" bson:"shipping_fee"`
}

type DeliveryArea struct {
	Name        string  `json:"name" bson:"name"`
	ShippingFee float64 `json:"shipping_fee,omitempty" bson:"shipping_fee,omitempty"`
}

type Purchase struct {
	Id                   entity.DatabaseId   `json:"id" bson:"_id"`
	ProductId            entity.DatabaseId   `json:"product_id" bson:"product_id"`
	OwnerId              entity.DatabaseId   `json:"owner_id" bson:"owner_id"`
	Reference            string              `json:"reference,omitempty" bson:"reference,omitempty"`
	Type                 timeline.Type       `json:"payment_type" bson:"payment_type"`
	Email                string              `json:"email" bson:"email"`
	Amount               float64             `json:"amount" bson:"amount"`
	ShippingFee          float64             `json:"shipping_fee" bson:"shipping_fee"`
	TimelineAmount       float64             `json:"timeline_amount" bson:"timeline_amount"`
	TotalAmount          float64             `json:"total_amount" bson:"total_amount"`
	NumberOfInstallments int                 `json:"number_of_installments,omitempty" bson:"number_of_installments,omitempty"`
	Frequency            Frequency           `json:"payment_frequency" bson:"payment_frequency"`
	DeliveryArea         string              `json:"delivery_area" bson:"delivery_area"`
	Successful           bool                `json:"successful" bson:"successful"`
	CreatedBy            entity.DatabaseId   `json:"created_by" bson:"created_by"`
	AdditionalDetails    string              `json:"additional_details,omitempty" bson:"additional_details,omitempty"`
	Timeline             []timeline.Timeline `json:"timeline" bson:"timeline"`
	Duration             time.Duration       `json:"-" bson:"-"`
	CreatedAt            time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at" bson:"updated_at"`
}

type Update struct {
	PurchaseId string
	Reference  string  `json:"reference"`
	Amount     float64 `json:"amount"`
}
