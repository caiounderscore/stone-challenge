package register

import "fmt"

type Register struct {
	Id            int
	CostumerItems []int
	Training      bool
}

func NewRegister(id int, t bool) *Register {
	return &Register{
		CostumerItems: make([]int, 0),
		Training:      t,
		Id:            id,
	}
}

func BuildRegisters(r []int) []*Register {
	var registers []*Register
	for i := 1; i <= r[0]; i++ {
		if i == r[0] {
			registers = append(registers, NewRegister(i, true))

		} else {
			registers = append(registers, NewRegister(i, false))
		}
	}
	return registers
}

func GetSmallRegister(rs []*Register) *Register {
	i := len(rs) - 1
	smallest := rs[i]
	for _, r := range rs {
		if len(r.CostumerItems) <= len(smallest.CostumerItems) {
			return r

		}
	}
	return smallest

}

func GetSmallRegisterItemRecord(rs []*Register) *Register {
	// i := rs[1]
	smallest := rs[1]
	for _, r := range rs {
		if r.IsEmpty() {
			return r
		}
		if smallest.IsEmpty() {
			return smallest
		}
		lastIdFewest := smallest.Size() - 1
		lastIdRegister := r.Size() - 1
		if r.CostumerItems[lastIdRegister] < smallest.CostumerItems[lastIdFewest] {
			smallest = r
		}
	}

	return smallest

}

func (r *Register) IsEmpty() bool {
	return len(r.CostumerItems) == 0
}

func (r *Register) Push(i int) {
	r.CostumerItems = append(r.CostumerItems, i)
}

func (r *Register) Size() int {
	return len(r.CostumerItems)
}

func (r *Register) ProcessItem(t *int) {
	if len(r.CostumerItems) == 0 {
		fmt.Println("eipa")
		return
	}
	for 0 == r.CostumerItems[0] {
		fmt.Println("OOOO")
		fmt.Println(r.CostumerItems[0])
		if r.CostumerItems[0] == 0 {
			fmt.Println("UEAP")
			continue
		}
		r.CostumerItems[0]--

		if r.Training {
			*t += 2
		} else {
			*t++
		}

	}

}

func (r *Register) pop() {
	if r.IsEmpty() {
		return
	} else {
		*&r.CostumerItems = (r.CostumerItems)[1:]
		return
	}
}
