package transfer

const bonusPercent int = 5

// Total - calculates weight of someones dog
func Total(amount int) (bonus int) {
	if amount >= 0 {
		bonus = amount + amount*bonusPercent/1000
		return
	}
	return 0
}
