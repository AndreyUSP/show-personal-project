package bank

import (
	"errors"
	"show-personal-project/internal/card"
)

type Bank struct {
	cards map[string]*card.Card
}

func NewBank() *Bank {
	return &Bank{
		cards: map[string]*card.Card{
			"1234": card.NewCard("1234", "0000", 1000.0),
		},
	}
}

func (b *Bank) VerifyPIN(c *card.Card, pin string) bool {
	return c.PIN == pin
}

func (b *Bank) Withdraw(c *card.Card, amount float64) error {
	if c.Balance < amount {
		return errors.New("Недостаточно средств")
	}
	c.Balance -= amount
	return nil
}

func (b *Bank) Deposit(c *card.Card, amount float64) {
	c.Balance += amount
}

func (b *Bank) Balance(c *card.Card) float64 {
	return c.Balance
}
