package types

// Money ...
type Money int64

// Currency ...
type Currency string

// TJS ...
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN ...
type PAN string

// Card ...
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment представляет информацию о платеже.
type Payment struct {
	ID     int
	Amount Money
}

// PaymentSource ...
type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
