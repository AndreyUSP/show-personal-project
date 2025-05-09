// internal/card/card.go
package card

type Card struct {
	Number  string
	PIN     string
	Balance float64
}

func NewCard(number, pin string, balance float64) *Card {
	return &Card{Number: number, PIN: pin, Balance: balance}
}
