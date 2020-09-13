package product

import (
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/string_utilities"
	"strings"
	"time"
)

func New() *Product {
	newProduct := new(Product)
	newProduct.Id = entity.NewDatabaseId()
	newProduct.ProductId = entity.NewDefaultId().String()

	currentTime := time.Now()
	newProduct.CreatedAt = currentTime
	newProduct.UpdatedAt = currentTime

	return newProduct
}

func (product *Product) CanBeCreated() *response.BaseResponse {
	if string_utilities.IsEmpty(strings.TrimSpace(product.Name)) {
		return response.NewBadRequestError("invalid product name")
	}
	if product.Amount == float64(0) {
		return response.NewBadRequestError("invalid product price")
	}

	if product.AllowInstallment {
		if product.MaxInstallment <= 1 {
			return response.NewBadRequestError("max installment must be greater than 1")
		}
		if len(product.PaymentFrequencies) == 0 {
			return response.NewBadRequestError("invalid payment frequencies")
		} else {
			for _, frequency := range product.PaymentFrequencies {
				if !frequency.IsValidFrequency() {
					message := fmt.Sprintf("%s is not a valid payment frequency", frequency)
					return response.NewBadRequestError(message)
				}
			}
		}
	} else {
		if len(product.PaymentFrequencies) != 0 {
			return response.NewBadRequestError("payment frequencies not allowed")
		}
		product.MaxInstallment = 1
	}
	if len(product.DeliveryAreas) == 0 {
		return response.NewBadRequestError("delivery area is empty")
	}
	return nil
}

func (product *Product) CanBePurchased(userId string, purchase *purchase.Purchase) *response.BaseResponse {
	if !purchase.HasValidPaymentFrequency() {
		return response.NewBadRequestError("invalid payment frequency")
	}
	if purchase.Type == timeline.TypeInstallment && !product.AllowInstallment {
		return response.NewBadRequestError("product does not allow installment payment")
	}
	if userId == product.OwnerId.Hex() {
		return response.NewBadRequestError("you can't buy your own product")
	}
	if purchase.NumberOfInstallments > product.MaxInstallment {
		return response.NewBadRequestError("specify a lower payment installment number")
	}
	isValidArea := false
	for _, area := range product.DeliveryAreas {
		if purchase.DeliveryArea == area {
			isValidArea = true
			break
		}
	}
	if !isValidArea {
		return response.NewBadRequestError("specified delivery area is not valid for product")
	}
	return nil
}
