package main

import (
	"fmt"
	"os"

	"github.com/caiounderscore/stone-challenge/costumer"
	"github.com/caiounderscore/stone-challenge/file"
	"github.com/caiounderscore/stone-challenge/register"
)

var registers []register.Register

func main() {
	// var wg sync.WaitGroup

	pathFile := os.Args[1]
	rawRegisters, rawCostumers := file.ParseFile(pathFile)
	// fmt.Println(rawRegisters, rawCostumers)
	costumer := costumer.BuildCostumers(rawCostumers)
	// fmt.Println(costumers)
	registers := register.BuildRegisters(rawRegisters)
	// fmt.Println(rawRegisters)
	fmt.Println(registers)
	// var tchan chan int
	// tchan = make(chan int, 1)
	var t = 1
	for _, c := range costumer {

		for _, r := range registers {
			r.ProcessItem(&t)
		}

		c.SelectRegister(registers)
		// fmt.Println(c.Type)
		c.Register.Push(c.Items)
		fmt.Println(c.Register)

	}

	fmt.Println(t)
	// wg.Wait()

	// close(tchan)
	// // go func(ch chan int, wg *sync.WaitGroup) {
	// fmt.Println(<-tchan)

	// }(tchan, &wg)
	// }(&wg)

}

// func choiceRegister(c, costumer.Costumer, rs []register.Register) register.Register {
// 	for _, r := range rs {
// 		if r.IsEmpty() {
// 			return r
// 		} else if c.Type == "B" && r.Costumers {

// 		}
// 	}
// }
