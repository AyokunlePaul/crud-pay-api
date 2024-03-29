package product

import (
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/string_utilities"
	"github.com/thoas/go-funk"
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
	} else {
		product.MaxInstallment = 0
	}
	for _, frequency := range product.PaymentFrequencies {
		if !frequency.IsValidFrequency() {
			message := fmt.Sprintf("%s is not a valid payment frequency", frequency)
			return response.NewBadRequestError(message)
		}
	}
	if len(product.DeliveryAreas) == 0 {
		return response.NewBadRequestError("delivery area is empty")
	}
	if len(product.DeliveryGroups) != 0 {
		for _, group := range product.DeliveryGroups {
			for areaIndex, area := range product.DeliveryAreas {
				if funk.Contains(group.Areas, area.Name) {
					//This will update the shipping fee of the area with that of it's
					//designated group
					product.DeliveryAreas[areaIndex].ShippingFee = group.ShippingFee
				}
			}
		}
	}
	return nil
}

func (product *Product) CanBeUpdatedWith(newProduct *Product) *response.BaseResponse {
	if newProduct.AllowInstallment {
		if product.MaxInstallment <= 1 {
			return response.NewBadRequestError("max installment must be greater than 1")
		}
		product.MaxInstallment = newProduct.MaxInstallment
		product.AllowInstallment = newProduct.AllowInstallment
	} else {
		product.MaxInstallment = 0
		product.AllowInstallment = false
	}
	if len(newProduct.PaymentFrequencies) != 0 {
		for _, frequency := range newProduct.PaymentFrequencies {
			if !frequency.IsValidFrequency() {
				message := fmt.Sprintf("%s is not a valid payment frequency", frequency)
				return response.NewBadRequestError(message)
			}
		}
		product.PaymentFrequencies = newProduct.PaymentFrequencies
	}
	if !string_utilities.IsEmpty(newProduct.Name) {
		product.Name = newProduct.Name
	}
	if newProduct.Amount != float64(0) {
		product.Amount = newProduct.Amount
	}
	if len(newProduct.Pictures) != 0 {
		product.Pictures = newProduct.Pictures
	}
	if len(newProduct.DeliveryAreas) != 0 {
		if len(newProduct.DeliveryGroups) != 0 {
			for _, group := range newProduct.DeliveryGroups {
				for _, area := range newProduct.DeliveryAreas {
					if funk.Contains(group.Areas, area) {
						//This will update the shipping fee of the area with that of it's
						//designated group
						area.ShippingFee = group.ShippingFee
					}
				}
			}
			product.DeliveryGroups = append(product.DeliveryGroups, newProduct.DeliveryGroups...)
		}
		product.DeliveryAreas = append(product.DeliveryAreas, newProduct.DeliveryAreas...)
	}
	return nil
}

func (product *Product) CanBePurchased(userId string, purchase *purchase.Purchase) *response.BaseResponse {
	if !string_utilities.IsValidEmail(purchase.Email) {
		return response.NewBadRequestError("invalid email address")
	}
	if !purchase.Type.IsValidPaymentType() {
		return response.NewBadRequestError("invalid payment type")
	}
	if purchase.Type == timeline.TypeInstallment && !product.AllowInstallment {
		return response.NewBadRequestError("product does not allow installment")
	}
	if purchase.Type == timeline.TypeRecurring && len(product.PaymentFrequencies) == 0 {
		return response.NewBadRequestError("product does not allow recurring payment")
	}

	isValidPaymentFrequency := false
	for _, frequency := range product.PaymentFrequencies {
		if frequency == purchase.Frequency {
			isValidPaymentFrequency = true
			break
		}
	}
	if (purchase.Type == timeline.TypeInstallment || purchase.Type == timeline.TypeRecurring) && !isValidPaymentFrequency {
		return response.NewBadRequestError("invalid payment frequency")
	}

	if userId == product.OwnerId.Hex() {
		return response.NewBadRequestError("you can't buy your own product")
	}
	if purchase.NumberOfInstallments <= 0 || purchase.NumberOfInstallments > product.MaxInstallment {
		return response.NewBadRequestError("invalid number of installment")
	}
	isValidArea := false
	for _, area := range product.DeliveryAreas {
		if purchase.DeliveryArea == area.Name {
			isValidArea = true
			purchase.ShippingFee = area.ShippingFee
			break
		}
	}
	if !isValidArea {
		return response.NewBadRequestError("invalid delivery area")
	}

	purchase.UpdatePaymentDuration()
	return nil
}
