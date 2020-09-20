package purchase_payload

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
)

type PurchasePayload struct {
	Payload map[string]interface{}
}

func (payload *PurchasePayload) ToDomain() *purchase.Purchase {
	domainPurchase := purchase.New()
	if productId, ok := payload.Payload["product_id"].(string); ok {
		domainPurchase.ProductId, _ = entity.StringToCrudPayId(productId)
	}
	if paymentType, ok := payload.Payload["payment_type"].(string); ok {
		domainPurchase.Type = timeline.Type(paymentType)
	}
	if email, ok := payload.Payload["email"].(string); ok {
		domainPurchase.Email = email
	}
	if numberOfInstallments, ok := payload.Payload["number_of_installments"].(int); ok {
		domainPurchase.NumberOfInstallments = numberOfInstallments
	}

	if additionalDetails, ok := payload.Payload["additional_details"].(string); ok {
		domainPurchase.AdditionalDetails = additionalDetails
	}
	return domainPurchase
}
