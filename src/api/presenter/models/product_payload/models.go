package product_payload

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
)

type ProductPayload struct {
	Payload map[string]interface{}
}

func (payload *ProductPayload) ToDomain() *product.Product {
	domainProduct := product.New()
	if productName, ok := payload.Payload["product_name"].(string); ok {
		domainProduct.Name = productName
	}
	if allowInstallment, ok := payload.Payload["allow_installment"].(bool); ok {
		domainProduct.AllowInstallment = allowInstallment
	}
	if pictures, ok := payload.Payload["pictures"].([]string); ok {
		domainProduct.Pictures = pictures
	}
	if maxInstallment, ok := payload.Payload["max_installment"].(float64); ok {
		domainProduct.MaxInstallment = int(maxInstallment)
	}
	if price, ok := payload.Payload["price"].(float64); ok {
		domainProduct.Amount = price
	}
	//if paymentFrequencies, ok := payload.Payload["payment_frequency"].([]string); ok {
	//	domainProduct.PaymentFrequencies = paymentFrequencies
	//}
	return domainProduct
}
