package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"time"
)

func NewPaymentTimeline(product product.Product, purchase Purchase) []Timeline {
	totalAmount := product.Amount
	numberOfInstallments := int(purchase.NumberOfInstallments)

	switch purchase.Type {
	case TypeOneTime:
		return []Timeline{{
			Id:                  entity.NewCrudPayId(),
			Paid:                false,
			Amount:              totalAmount,
			ExpectedPaymentDate: time.Now(),
		}}
	case TypeInstallment:
		timelines := make([]Timeline, purchase.NumberOfInstallments)
		amountPerTimeline := totalAmount / float64(numberOfInstallments)

		lastPaymentMade := time.Now()
		for i := 0; i < numberOfInstallments; i++ {
			currentTimeline := Timeline{
				Id:                  entity.NewCrudPayId(),
				Paid:                false,
				Amount:              amountPerTimeline,
				ExpectedPaymentDate: lastPaymentMade,
			}
			timelines = append(timelines, currentTimeline)
			lastPaymentMade = lastPaymentMade.Add(purchase.Duration)
		}

		return timelines
	default:
		return []Timeline{{
			Id:                  entity.NewCrudPayId(),
			Paid:                false,
			Amount:              totalAmount,
			ExpectedPaymentDate: time.Now(),
		}}
	}
}

func (purchase *Purchase) HasValidPaymentFrequency() bool {
	isValidFrequency := purchase.Frequency.IsValidFrequency()
	if !isValidFrequency {
		return false
	}
	switch purchase.Frequency {
	case BiWeekly:
		purchase.Duration = DurationBiWeekly
	case Monthly:
		purchase.Duration = DurationMonthly
	case Quarterly:
		purchase.Duration = DurationQuarterly
	case BiAnnually:
		purchase.Duration = DurationBiAnnually
	default:
		purchase.Duration = DurationAnnually
	}
	return true
}

func (frequency PaymentFrequency) IsValidFrequency() bool {
	return frequency == BiWeekly || frequency == Monthly ||
		frequency == Quarterly || frequency == BiAnnually || frequency == Annually
}
