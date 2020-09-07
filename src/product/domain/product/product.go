package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product/delivery"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities/string_utilities"
	"github.com/myesui/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Product struct {
	Id               primitive.ObjectID `json:"id" bson:"_id"`
	Name             string             `json:"product_name" bson:"product_name"`
	ProductId        string             `json:"product_id" bson:"product_id"`
	AllowInstallment bool               `json:"allow_installment,omitempty" bson:"allow_installment"`
	Pictures         []string           `json:"pictures,omitempty" bson:"pictures"`
	MaxInstallments  int64              `json:"max_installments,omitempty" bson:"max_installments"`
	Price            float64            `json:"price,omitempty" bson:"price"`
	DeliveryGroups   []delivery.Group   `json:"delivery_groups,omitempty" bson:"delivery_groups,omitempty"`
	OwnerId          primitive.ObjectID `json:"owner_id,omitempty" bson:"owner_id"`
}

func (product *Product) IsValidProduct() *response.BaseResponse {
	if string_utilities.IsEmpty(strings.TrimSpace(product.Name)) {
		return response.NewBadRequestError("invalid product name")
	}
	if product.Price <= 0.0 {
		return response.NewBadRequestError("invalid product price")
	}
	product.ProductId = uuid.NewV4().String()
	return nil
}
