package card

import (
	"bank/pkg/bank/types"
)

const (
	depositLimit  types.Money = 50000_00
	withdrawLimit types.Money = 20000_00
	bonusLimit    types.Money = 5000_00
)

// IssueCard ...
func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Name:     name,
		Color:    color,
		Active:   true,
	}
}

// Withdraw ...
func Withdraw(card types.Card, amount types.Money) types.Card {
	if card.Active && amount > 0 && amount < card.Balance && amount <= withdrawLimit {
		card.Balance = card.Balance - amount
	}
	return card
}

// Deposit ...
func Deposit(card *types.Card, amount types.Money) {
	if !(*card).Active {
		return
	}
	if amount > depositLimit {
		return
	}
	if amount < 0 {
		return
	}

	card.Balance += amount
}

// AddBonus ...
func AddBonus(card *types.Card, percent, daysInMonth, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance < 0 {
		return
	}
	if card.MinBalance < 0 {
		return
	}
	bonus := types.Money((int(card.MinBalance) * percent * daysInMonth / daysInYear) / 100)
	if bonus > bonusLimit {
		return
	}
	card.Balance += bonus
}

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(cards []types.Card) types.Money {
	var totalAmount types.Money = 0
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance < 0 {
			continue
		}
		totalAmount += card.Balance
	}
	return totalAmount
}
