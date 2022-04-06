package grocery

// import (
// 	"github.com/caiounderscore/stone-challenge/costumer"
// 	"github.com/caiounderscore/stone-challenge/register"
// )

// type Grocery struct {
// 	Registers []*register.Register
// 	Costumers []*costumer.Costumer
// }

// func (g *Grocery) GetSmallRegister() register.Register {
// 	i := len(g.Registers)
// 	smallest := g.Registers[i]
// 	for _, r := range g.Registers {
// 		if len(r.CostumerItems) < len(smallest.CostumerItems) {
// 			smallest = r
// 		}
// 	}
// 	return *smallest

// }

// func (g *Grocery) GetSmallRegisterItemRecord() register.Register {
// 	i := len(g.Registers)
// 	smallest := g.Registers[i]
// 	for _, r := range g.Registers {
// 		if len(r.CostumerItems) == 0 {
// 			return *r
// 		}

// 		lastIdFewest := len(smallest.CostumerItems) - 1
// 		lastIdRegister := len(r.CostumerItems) - 1

// 		if r.CostumerItems[lastIdRegister] < smallest.CostumerItems[lastIdFewest] {
// 			smallest = r
// 		}
// 	}
// 	return *smallest

// }
