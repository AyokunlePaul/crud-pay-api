package product_payload

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
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
	if maxInstallments, ok := payload.Payload["max_installments"].(int); ok {
		domainProduct.MaxInstallments = maxInstallments
	}
	if price, ok := payload.Payload["price"].(float64); ok {
		domainProduct.Amount = price
	}
	if paymentFrequencies, ok := payload.Payload["payment_frequency"].([]purchase.PaymentFrequency); ok {
		domainProduct.PaymentFrequencies = paymentFrequencies
	}
	return domainProduct
}
