package deposit

//Calculate - calculates some weird staff
func Calculate(amount int, currency string) (min, max int) {
	minPercent, maxPercent := percentFor(currency)
	min = (amount * minPercent / 100) / 12
	max = (amount * maxPercent / 100) / 12
	return
}

func percentFor(currency string) (min, max int) {
	switch currency {
	case "TJS":
		min = 4
		max = 6
	case "USD":
		min = 1
		max = 2
	default:
		min = 0
		max = 0
	}
	return
}
