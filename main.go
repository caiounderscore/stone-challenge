package main

import (
	"fmt"
	"os"

	"github.com/caiounderscore/stone-challenge/file"
	"github.com/caiounderscore/stone-challenge/grocery"
)

var registers []grocery.Register

func main() {
	pathFile := os.Args[1]
	rawRegisters, rawCustomers := file.ParseFile(pathFile)
	Customers := grocery.BuildCustomers(rawCustomers)
	registers := grocery.BuildRegisters(rawRegisters)
	var t = 1
	gr := grocery.NewGrocery(registers, Customers)
	gr.SortCustomers()
	for {
		gr.ProcessItem(t)
		for _, c := range gr.Customers {
			if c.ArrivalTime == t {
				c.SelectRegister(registers, t)
				c.Register.Push(*c)

			}
		}

		if grocery.CheckRegistersIfEmpty(registers) {
			break
		}
		t++

	}
	fmt.Printf("Finished at: t=%v minutes\n", t)

}
