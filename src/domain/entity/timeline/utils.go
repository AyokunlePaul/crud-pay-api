package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"math"
	"time"
)

func NewTimeline(
	purchaseId entity.DatabaseId, amount, shippingFee float64,
	numberOfInstallments int, duration time.Duration, purchaseType Type,
) []Timeline {
	switch purchaseType {
	case TypeOneTime:
		return []Timeline{{
			Id:                  entity.NewDatabaseId(),
			PurchaseId:          purchaseId,
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	case TypeInstallment:
		timelines := make([]Timeline, numberOfInstallments)
		amountPerTimeline := amount / float64(numberOfInstallments)
		amountPerTimeline = math.Ceil(amountPerTimeline*100) / 100 //Round it off to 2 decimal places if needed

		lastPaymentMade := time.Now()
		for index := 0; index < numberOfInstallments; index++ {
			currentTimeline := Timeline{
				Id:                  entity.NewDatabaseId(),
				PurchaseId:          purchaseId,
				Paid:                false,
				Amount:              amountPerTimeline,
				ExpectedPaymentDate: lastPaymentMade,
			}
			if index == 0 {
				currentTimeline.ShippingFee = shippingFee
			}
			timelines[index] = currentTimeline
			lastPaymentMade = lastPaymentMade.Add(duration)
		}
		return timelines
	default:
		//By default, we create 4 payment timeline for recurring payments
		var timelines []Timeline
		lastPaymentMade := time.Now()
		for index := 0; index < 4; index++ {
			currentTimeline := Timeline{
				Id:                  entity.NewDatabaseId(),
				Paid:                false,
				Amount:              amount,
				PurchaseId:          purchaseId,
				ShippingFee:         shippingFee,
				ExpectedPaymentDate: lastPaymentMade,
			}
			timelines = append(timelines, currentTimeline)
			lastPaymentMade = lastPaymentMade.Add(duration)
		}
		return timelines
	}
}

func (paymentType Type) IsValidPaymentType() bool {
	return paymentType == TypeOneTime || paymentType == TypeInstallment || paymentType == TypeRecurring
}
