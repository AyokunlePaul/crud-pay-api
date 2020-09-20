package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"math"
	"time"
)

func NewTimeline(
	purchaseId entity.DatabaseId, amount float64,
	numberOfInstallments int, duration time.Duration, purchaseType Type,
) []interface{} {
	switch purchaseType {
	case TypeOneTime:
		return []interface{}{Timeline{
			Id:                  entity.NewDatabaseId(),
			PurchaseId:          purchaseId,
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	case TypeInstallment:
		timelines := make([]interface{}, numberOfInstallments)
		amountPerTimeline := amount / float64(numberOfInstallments)
		amountPerTimeline = math.Ceil(amountPerTimeline*100) / 100

		lastPaymentMade := time.Now()
		for index := 0; index < numberOfInstallments; index++ {
			currentTimeline := Timeline{
				Id:                  entity.NewDatabaseId(),
				PurchaseId:          purchaseId,
				Paid:                false,
				Amount:              amountPerTimeline,
				ExpectedPaymentDate: lastPaymentMade,
			}
			timelines[index] = currentTimeline
			lastPaymentMade = lastPaymentMade.Add(duration)
		}
		return timelines
	default:
		return []interface{}{Timeline{
			Id:                  entity.NewDatabaseId(),
			PurchaseId:          purchaseId,
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	}
}

func (paymentType Type) IsValidPaymentType() bool {
	return paymentType == TypeOneTime || paymentType == TypeInstallment || paymentType == TypeRecurring
}
