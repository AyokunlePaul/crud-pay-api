package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"time"
)

func New() *Purchase {
	newPurchase := new(Purchase)
	newPurchase.Id = entity.NewDatabaseId()

	currentTime := time.Now()

	newPurchase.CreatedAt = currentTime
	newPurchase.UpdatedAt = currentTime

	return newPurchase
}

func (purchase *Purchase) UpdatePaymentDuration() {
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
}

func (frequency Frequency) IsValidFrequency() bool {
	return frequency == BiWeekly || frequency == Monthly ||
		frequency == Quarterly || frequency == BiAnnually || frequency == Annually
}
