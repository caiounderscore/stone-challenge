package costumer

import (
	"strconv"

	"github.com/caiounderscore/stone-challenge/register"
)

type Costumer struct {
	Id       int
	Type     string
	Items    int
	Register *register.Register
}

var Costumers []Costumer

func NewCostumer() Costumer {
	return Costumer{}
}

func (c *Costumer) SelectRegister(rs []*register.Register) {
	switch ctype := c.Type; ctype {
	case "A":
		r := register.GetSmallRegister(rs)
		c.Register = r
	case "B":
		r := register.GetSmallRegisterItemRecord(rs)
		c.Register = r
	default:
		panic("Invalid customer type")
	}
}

func BuildCostumers(rawCostumers [][]string) []Costumer {
	for _, v := range rawCostumers {
		costumer := NewCostumer()
		for i, v := range v {
			switch i {
			case 0:
				costumer.Type = v
			case 1:
				id, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				costumer.Id = id
			case 2:
				items, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				costumer.Items = items

			default:
				panic("Something is wrong")
			}
		}
		Costumers = append(Costumers, costumer)

	}
	return Costumers
}
