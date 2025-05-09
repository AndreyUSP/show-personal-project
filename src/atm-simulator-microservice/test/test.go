package test

import (
	"show-personal-project/internal/atm"
	"show-personal-project/internal/bank"
	"show-personal-project/internal/card"
	"testing"
)

func TestWithdraw(t *testing.T) {
	b := bank.NewBank()
	a := atm.NewATM(b)
	c := card.NewCard("1234", "0000", 1000)
	a.InsertCard(c)
	if !a.EnterPIN("0000") {
		t.Fatal("Не удалось ввести правильный PIN")
	}
	a.Withdraw(200)
	if b.Balance(c) != 800 {
		t.Fatalf("Ожидалось 800, получено %.2f", b.Balance(c))
	}
}
