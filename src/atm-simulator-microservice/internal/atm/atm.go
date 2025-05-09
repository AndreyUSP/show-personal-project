// internal/atm/atm.go
package atm

import (
	"fmt"
	"show-personal-project/internal/bank"
	"show-personal-project/internal/card"
)

type ATM struct {
	bank       *bank.Bank
	inserted   *card.Card
	pinEntered bool
}

func NewATM(b *bank.Bank) *ATM {
	return &ATM{bank: b}
}

func (a *ATM) InsertCard(c *card.Card) {
	a.inserted = c
	a.pinEntered = false
	fmt.Println("Карта вставлена.")
}

func (a *ATM) EnterPIN(pin string) bool {
	if a.inserted == nil {
		fmt.Println("Нет вставленной карты.")
		return false
	}
	if a.bank.VerifyPIN(a.inserted, pin) {
		a.pinEntered = true
		fmt.Println("PIN принят.")
		return true
	} else {
		fmt.Println("Неверный PIN.")
		return false
	}
}

func (a *ATM) Withdraw(amount float64) {
	if !a.isAuthorized() {
		return
	}
	err := a.bank.Withdraw(a.inserted, amount)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println("Снятие выполнено.")
	}
}

func (a *ATM) Deposit(amount float64) {
	if !a.isAuthorized() {
		return
	}
	a.bank.Deposit(a.inserted, amount)
	fmt.Println("Пополнение выполнено.")
}

func (a *ATM) CheckBalance() {
	if !a.isAuthorized() {
		return
	}
	fmt.Printf("Баланс: %.2f\n", a.bank.Balance(a.inserted))
}

func (a *ATM) EjectCard() {
	a.inserted = nil
	a.pinEntered = false
	fmt.Println("Карта извлечена.")
}

func (a *ATM) isAuthorized() bool {
	if a.inserted == nil {
		fmt.Println("Нет карты.")
		return false
	}
	if !a.pinEntered {
		fmt.Println("PIN не введён.")
		return false
	}
	return true
}
