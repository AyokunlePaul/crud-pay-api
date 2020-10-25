package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"time"
)

type Product struct {
	Id                 entity.DatabaseId       `json:"id" bson:"_id"`
	Name               string                  `json:"product_name" bson:"product_name"`
	ProductId          string                  `json:"-" bson:"product_id"`
	AllowInstallment   bool                    `json:"allow_installment" bson:"allow_installment"`
	PaymentFrequencies []purchase.Frequency    `json:"payment_frequency,omitempty" bson:"payment_frequency,omitempty"`
	Pictures           []string                `json:"pictures,omitempty" bson:"pictures,omitempty"`
	MaxInstallment     int                     `json:"max_installment,omitempty" bson:"max_installment,omitempty"`
	IsDeleted          bool                    `json:"-" bson:"is_deleted"`
	Amount             float64                 `json:"amount" bson:"amount"`
	DeliveryGroups     []purchase.Group        `json:"delivery_groups,omitempty" bson:"delivery_groups"`
	DeliveryAreas      []purchase.DeliveryArea `json:"delivery_areas,omitempty" bson:"delivery_areas,omitempty"`
	OwnerId            entity.DatabaseId       `json:"owner_id" bson:"owner_id"`
	CreatedAt          time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at" bson:"updated_at"`
}
