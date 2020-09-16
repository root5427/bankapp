package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20000_00, Active: true}, 10000_00)
	fmt.Println(result.Balance)
	// Output:
	// 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 1000, Active: true}, 1500)
	fmt.Println(result.Balance)
	// Output:
	// 1000
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 1000, Active: false}, 1500)
	fmt.Println(result.Balance)
	// Output:
	// 1000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 600000_00, Active: true}, 21000_00)
	fmt.Println(result.Balance)
	// Output:
	// 60000000
}

func ExampleDeposit_normal() {
	card := types.Card{
		Balance: 20_000_00,
		Active:  true,
	}
	var amount types.Money = 10_000_00
	Deposit(&card, amount)
	fmt.Println(card.Balance)

	// Output:
	// 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{
		Balance: 20_000_00,
		Active:  false,
	}
	var amount types.Money = 10_000_00
	Deposit(&card, amount)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleDeposit_overDepositLimit() {
	card := types.Card{
		Balance: 20_000_00,
		Active:  false,
	}
	var amount types.Money = 50_000_01
	Deposit(&card, amount)
	fmt.Println(card.Balance)

	// Output:
	// 2000000
}

func ExampleDeposit_negativeBalance() {
	card := types.Card{
		Balance: -20_000_00,
		Active:  true,
	}
	var amount types.Money = 50_000_00
	Deposit(&card, amount)
	fmt.Println(card.Balance)

	// Output:
	// 3000000
}

func ExampleAddBonus_positive() {
	card := types.Card{
		Balance:    20_000_00,
		MinBalance: 10_000_00,
		Active:     true,
	}
	oldBalance := card.Balance
	AddBonus(&card, 3, 30, 365)
	newBalance := card.Balance
	fmt.Println(newBalance - oldBalance)
	// Output:
	// 2465
}

func ExampleAddBonus_inactive() {
	card := types.Card{
		Balance:    20_000_00,
		MinBalance: 10_000_00,
		Active:     false,
	}
	oldBalance := card.Balance
	AddBonus(&card, 3, 30, 365)
	newBalance := card.Balance
	fmt.Println(newBalance - oldBalance)
	// Output:
	// 0
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{
		Balance:    -20_000_00,
		MinBalance: 10_000_00,
		Active:     true,
	}
	oldBalance := card.Balance
	AddBonus(&card, 3, 30, 365)
	newBalance := card.Balance
	fmt.Println(newBalance - oldBalance)
	// Output:
	// 0
}

func ExampleAddBonus_overBonusLimit() {
	card := types.Card{
		Balance:    20_000_00,
		MinBalance: 25_000_00,
		Active:     true,
	}
	oldBalance := card.Balance
	AddBonus(&card, 3, 30, 365)
	newBalance := card.Balance
	fmt.Println(newBalance - oldBalance)
	// Output:
	// 6164
}

func ExampleAddBonus_equalBonusLimit() {
	card := types.Card{
		Balance:    20_000_00,
		MinBalance: 20_277_78,
		Active:     true,
	}
	oldBalance := card.Balance
	AddBonus(&card, 3, 30, 365)
	newBalance := card.Balance
	fmt.Println(newBalance - oldBalance)
	// Output:
	// 5000
}

func ExampleTotal_positive() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 3000000
}

func ExampleTotal_negativeBalance() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 2000000
}

func ExampleTotal_inactive() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  false,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	// Output: 0
}
