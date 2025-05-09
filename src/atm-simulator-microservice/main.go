// cmd/main.go
package main

import (
	"show-personal-project/internal/atm"
	"show-personal-project/internal/bank"
	"show-personal-project/internal/ui"
)

func main() {
	b := bank.NewBank()
	a := atm.NewATM(b)
	ui.Run(a)
}
