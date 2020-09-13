package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

func NewTimeline(
	purchaseId entity.DatabaseId, amount float64,
	numberOfInstallments int, duration time.Duration, purchaseType Type,
) []interface{} {
	switch purchaseType {
	case TypeOneTime:
		return []interface{}{Timeline{
			Id:                  entity.NewCrudPayId(),
			PurchaseId:          purchaseId,
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	case TypeInstallment:
		timelines := make([]interface{}, numberOfInstallments)
		amountPerTimeline := amount / float64(numberOfInstallments)

		lastPaymentMade := time.Now()
		for i := 0; i < numberOfInstallments; i++ {
			currentTimeline := Timeline{
				Id:                  entity.NewCrudPayId(),
				PurchaseId:          purchaseId,
				Paid:                false,
				Amount:              amountPerTimeline,
				ExpectedPaymentDate: lastPaymentMade,
			}
			timelines = append(timelines, currentTimeline)
			lastPaymentMade = lastPaymentMade.Add(duration)
		}

		return timelines
	default:
		return []interface{}{Timeline{
			Id:                  entity.NewCrudPayId(),
			PurchaseId:          purchaseId,
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	}
}
