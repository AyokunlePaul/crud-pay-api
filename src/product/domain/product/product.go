package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product/delivery"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	ProductName      string             `json:"product_name" bson:"product_name"`
	ProductId        string             `json:"product_id" bson:"product_id"`
	AllowInstallment bool               `json:"allow_installment" bson:"allow_installment"`
	Pictures         []string           `json:"pictures" bson:"pictures"`
	MaxInstallments  int64              `json:"max_installments" bson:"max_installments"`
	Price            float64            `json:"price" bson:"price"`
	DeliveryGroups   []delivery.Group   `json:"delivery_groups,omitempty" bson:"delivery_groups,omitempty"`
}
