package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"time"
)

type Product struct {
	Id                 entity.DatabaseId       `json:"id,omitempty" bson:"_id"`
	Name               string                  `json:"product_name,omitempty" bson:"product_name"`
	ProductId          string                  `json:"product_id,omitempty" bson:"product_id"`
	AllowInstallment   bool                    `json:"allow_installment,omitempty" bson:"allow_installment"`
	PaymentFrequencies []purchase.Frequency    `json:"payment_frequency,omitempty" bson:"payment_frequency"`
	Pictures           []string                `json:"pictures,omitempty" bson:"pictures"`
	MaxInstallment     int                     `json:"max_installment,omitempty" bson:"max_installment"`
	Amount             float64                 `json:"amount,omitempty" bson:"amount"`
	DeliveryGroups     []Group                 `json:"delivery_groups,omitempty" bson:"delivery_groups"`
	DeliveryAreas      []purchase.DeliveryArea `json:"delivery_areas" bson:"delivery_areas"`
	OwnerId            entity.DatabaseId       `json:"owner_id,omitempty" bson:"owner_id"`
	CreatedAt          time.Time               `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at,omitempty" bson:"updated_at"`
}

type Group struct {
	GroupId     entity.DatabaseId `json:"group_id" bson:"group_id"`
	Name        string            `json:"group_name" bson:"group_name"`
	Locations   []string          `json:"locations" bson:"locations"`
	ShippingFee float64           `json:"shipping_fee" bson:"shipping_fee"`
}
