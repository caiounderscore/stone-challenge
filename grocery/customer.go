package grocery

import (
	"strconv"
)

type Customer struct {
	Id            int
	ArrivalTime   int
	Type          string
	Items         int
	Register      *Register
	TimeToProcess int
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c *Customer) SelectRegister(rs []*Register, t int) {
	switch ctype := c.Type; ctype {
	case "A":
		r := getSmallRegister(rs)
		c.Register = r
	case "B":
		r := getSmallRegisterItemRecord(rs)
		c.Register = r
	default:
		panic("Invalid customer type")
	}

	if c.Register.Training {
		c.TimeToProcess = t + 2
	} else {
		c.TimeToProcess = t + 1
	}

}

func BuildCustomers(rawCustomers [][]string) []*Customer {
	for i, v := range rawCustomers {
		var id int = 1
		id += i
		Customer := NewCustomer()
		Customer.Id = id
		for i, v := range v {
			switch i {
			case 0:
				Customer.Type = v
			case 1:
				time, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				Customer.ArrivalTime = time
			case 2:
				items, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				Customer.Items = items

			default:
				panic("Something is wrong")
			}
		}
		Customers = append(Customers, Customer)

	}
	return Customers
}
