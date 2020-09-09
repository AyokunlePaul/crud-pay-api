package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

type Product struct {
	Id               entity.DatabaseId `json:"id,omitempty" bson:"_id"`
	Name             string            `json:"product_name,omitempty" bson:"product_name"`
	ProductId        string            `json:"product_id,omitempty" bson:"product_id"`
	AllowInstallment bool              `json:"allow_installment,omitempty" bson:"allow_installment"`
	Pictures         []string          `json:"pictures,omitempty" bson:"pictures"`
	MaxInstallments  int64             `json:"max_installments,omitempty" bson:"max_installments"`
	Price            float64           `json:"price,omitempty" bson:"price"`
	DeliveryGroups   []Group           `json:"delivery_groups,omitempty" bson:"delivery_groups"`
	OwnerId          entity.DatabaseId `json:"owner_id,omitempty" bson:"owner_id"`
	CreatedAt        time.Time         `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at,omitempty" bson:"updated_at"`
}

type Group struct {
	GroupId     entity.DatabaseId `json:"group_id" bson:"group_id"`
	Name        string            `json:"group_name" bson:"group_name"`
	Locations   []string          `json:"locations" bson:"locations"`
	ShippingFee float64           `json:"shipping_fee" bson:"shipping_fee"`
}
