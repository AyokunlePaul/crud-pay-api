package timeline

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

func NewPaymentTimeline(
	amount float64, numberOfInstallments int,
	duration time.Duration, purchaseType Type,
) []Timeline {
	switch purchaseType {
	case TypeOneTime:
		return []Timeline{{
			Id:                  entity.NewCrudPayId(),
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	case TypeInstallment:
		timelines := make([]Timeline, numberOfInstallments)
		amountPerTimeline := amount / float64(numberOfInstallments)

		lastPaymentMade := time.Now()
		for i := 0; i < numberOfInstallments; i++ {
			currentTimeline := Timeline{
				Id:                  entity.NewCrudPayId(),
				Paid:                false,
				Amount:              amountPerTimeline,
				ExpectedPaymentDate: lastPaymentMade,
			}
			timelines = append(timelines, currentTimeline)
			lastPaymentMade = lastPaymentMade.Add(duration)
		}

		return timelines
	default:
		return []Timeline{{
			Id:                  entity.NewCrudPayId(),
			Paid:                false,
			Amount:              amount,
			ExpectedPaymentDate: time.Now(),
		}}
	}
}
