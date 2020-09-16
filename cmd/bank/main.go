package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	card1 := types.Card{
		Balance:    0,
		MinBalance: 2_027_777_78,
		Active:     true,
	}
	oldBalance := card1.Balance
	card.AddBonus(&card1, 3, 30, 365)
	newBalance := card1.Balance
	fmt.Println(newBalance - oldBalance)

}
