package main

import (
	"fmt"
	"os"

	"github.com/caiounderscore/stone-challenge/file"
	"github.com/caiounderscore/stone-challenge/grocery"
)

var registers []grocery.Register

func main() {
	// var wg sync.WaitGroup

	pathFile := os.Args[1]
	rawRegisters, rawCustomers := file.ParseFile(pathFile)
	// fmt.Println(rawRegisters, rawCustomers)
	Customers := grocery.BuildCustomers(rawCustomers)
	// fmt.Println(Customers)
	registers := grocery.BuildRegisters(rawRegisters)
	// fmt.Println(rawRegisters)
	// fmt.Println(Customers)
	// var tchan chan int
	// tchan = make(chan int, 1)
	var t = 1
	gr := grocery.NewGrocery(registers, Customers)
	// var t2 = 1
	for {
		for _, c := range gr.Customers {
			if c.ArrivalTime == t {
				c.SelectRegister(registers, t)
				c.Register.Push(*c)
				gr.Customers = gr.Customers[1:]

			}
		}
		gr.ProcessItem(t)

		if grocery.CheckRegistersIfEmpty(registers) {
			break
		}
		t++

	}
	// i := len(gr.Registers)
	fmt.Println(gr.Registers[0].TimeToProcess)
	// wg.Wait()

	// close(tchan)
	// // go func(ch chan int, wg *sync.WaitGroup) {
	// fmt.Println(<-tchan)

	// }(tchan, &wg)
	// }(&wg)

}

// func choiceRegister(c, Customer.Customer, rs []register.Register) register.Register {
// 	for _, r := range rs {
// 		if r.IsEmpty() {
// 			return r
// 		} else if c.Type == "B" && r.Customers {

// 		}
// 	}
// }
