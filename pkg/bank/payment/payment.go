package payment

import (
	"bank/pkg/bank/types"
)

// Max возвращает максимальный платёж.
func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for _, payment := range payments {
		if payment.Amount > max.Amount {
			max = payment
		}
	}
	return max
}

// PaymentSources ...
func PaymentSources(cards []types.Card) []types.PaymentSource {
	var ps types.PaymentSource
	var pss []types.PaymentSource
	for _, card := range cards {
		if card.Balance < 0 {
			continue
		}
		if card.Active == false {
			continue
		}
		ps = types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		}
		pss = append(pss, ps)
	}
	return pss
}
