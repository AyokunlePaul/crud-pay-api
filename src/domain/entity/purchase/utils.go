package purchase

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
