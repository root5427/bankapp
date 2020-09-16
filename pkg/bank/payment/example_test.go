package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			Amount: 10_000_00,
		},
		{
			Amount: 20_000_00,
		},
		{
			Amount: 30_000_00,
		},
	}
	fmt.Println(Max(payments).Amount)
	// Output: 3000000
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "card1",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "card2",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			PAN:     "card3",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "card4",
			Balance: 10_000_00,
			Active:  false,
		},
	}

	pss := PaymentSources(cards)

	for _, ps := range pss {
		fmt.Println(ps.Number)
	}

	// Output:
	// card1
	// card3
}
