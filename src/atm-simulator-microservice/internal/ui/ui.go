// internal/ui/ui.go
package ui

import (
	"bufio"
	"fmt"
	"os"
	"show-personal-project/internal/atm"
	"show-personal-project/internal/card"
	"strconv"
)

func Run(a *atm.ATM) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1) Вставить карту\n2) Ввести PIN\n3) Баланс\n4) Снять\n5) Пополнить\n6) Извлечь карту\n0) Выход")
		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			a.InsertCard(card.NewCard("1234", "0000", 1000.0))
		case "2":
			fmt.Print("Введите PIN: ")
			scanner.Scan()
			a.EnterPIN(scanner.Text())
		case "3":
			a.CheckBalance()
		case "4":
			fmt.Print("Сумма снятия: ")
			scanner.Scan()
			amt, _ := strconv.ParseFloat(scanner.Text(), 64)
			a.Withdraw(amt)
		case "5":
			fmt.Print("Сумма пополнения: ")
			scanner.Scan()
			amt, _ := strconv.ParseFloat(scanner.Text(), 64)
			a.Deposit(amt)
		case "6":
			a.EjectCard()
		case "0":
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
