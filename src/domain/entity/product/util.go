package product

import (
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"time"
)

func New() *Product {
	newProduct := new(Product)
	newProduct.Id = entity.NewCrudPayId()
	newProduct.ProductId = entity.NewDefaultId().String()

	currentTime := time.Now()
	newProduct.CreatedAt = currentTime
	newProduct.UpdatedAt = currentTime
	newProduct.DeliveryGroups = []Group{}
	newProduct.Pictures = []string{}

	return newProduct
}

func (product *Product) CanBeCreated() *response.BaseResponse {
	if product.AllowInstallment {
		if product.MaxInstallments <= 1 {
			message := fmt.Sprintf("max installment must be greater than 1")
			return response.NewBadRequestError(message)
		}
		if len(product.PaymentFrequencies) == 0 {
			message := fmt.Sprintf("payment frequencies not specified")
			return response.NewBadRequestError(message)
		} else {
			for _, frequency := range product.PaymentFrequencies {
				if !frequency.IsValidFrequency() {
					message := fmt.Sprintf("%s is not a valid frequency", frequency)
					return response.NewBadRequestError(message)
				}
			}
		}
	} else {
		if len(product.PaymentFrequencies) != 0 {
			message := fmt.Sprintf("payment frequencies not allowed")
			return response.NewBadRequestError(message)
		}
	}
	return nil
}