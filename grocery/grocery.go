package grocery

import "sort"

type Grocery struct {
	Registers      []*Register
	Customers      []*Customer
	TotalCustomers int
}

func NewGrocery(rs []*Register, cs []*Customer) *Grocery {
	return &Grocery{
		Registers:      rs,
		Customers:      cs,
		TotalCustomers: len(cs),
	}
}

func (g *Grocery) SortCustomers() {
	sort.Slice(g.Customers, func(i, j int) bool {
		if g.Customers[i].Items == g.Customers[j].Items {
			return g.Customers[i].Type == "A"
		}

		return g.Customers[i].Items < g.Customers[j].Items
	})
}

func (g *Grocery) ProcessItem(t int) {
	for _, r := range g.Registers {
		if len(r.Customers) != 0 && r.Customers[0].TimeToProcess == t {
			r.Customers[0].Items--
			if r.Customers[0].Items <= 0 {
				r.Pop()
				if len(r.Customers) == 0 {
					continue
				}
			}

			if r.Training {
				r.Customers[0].TimeToProcess = t + 2
			} else {
				r.Customers[0].TimeToProcess = t + 1
			}
		}
	}
}
