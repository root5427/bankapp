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
