package product_payload

import (
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
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
	if maxInstallments, ok := payload.Payload["max_installments"].(int64); ok {
		domainProduct.MaxInstallments = maxInstallments
	}
	if price, ok := payload.Payload["price"].(float64); ok {
		domainProduct.Price = price
	}
	logger.Info(fmt.Sprintf("%v", domainProduct))
	return domainProduct
}
