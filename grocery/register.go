package grocery

type Register struct {
	Id            int
	Customers     []Customer
	Training      bool
	TimeToProcess int
}

func NewRegister(id int, t bool) *Register {
	return &Register{
		Customers: make([]Customer, 0),
		Training:  t,
		Id:        id,
	}
}

func BuildRegisters(r []int) []*Register {
	var Registers []*Register
	for i := 1; i <= r[0]; i++ {
		if i == r[0] {
			Registers = append(Registers, NewRegister(i, true))

		} else {
			Registers = append(Registers, NewRegister(i, false))
		}
	}
	return Registers
}

func getSmallRegister(rs []*Register) *Register {
	smallest := rs[0]
	for _, r := range rs {
		if r.Size() < smallest.Size() {
			smallest = r

		}
	}
	return smallest

}

func getSmallRegisterItemRecord(rs []*Register) *Register {
	smallest := rs[0]
	for _, r := range rs {
		if len(r.Customers) == 0 {
			return r
		}
		lastIdFewest := smallest.Size() - 1
		lastIdRegister := r.Size() - 1
		if r.Customers[lastIdRegister].Items < smallest.Customers[lastIdFewest].Items {
			return r
		}
	}

	return smallest

}

func CheckRegistersIfEmpty(rs []*Register) bool {
	if len(rs) == 1 {
		return rs[0].IsEmpty()
	}
	return rs[0].IsEmpty() && rs[1].IsEmpty()

}

func (r *Register) IsEmpty() bool {
	return r.Size() == 0
}

func (r *Register) Push(c Customer) {
	r.Customers = append(r.Customers, c)
}

func (r *Register) Size() int {
	return len(r.Customers)
}

func (r *Register) Pop() {
	if r.IsEmpty() {
		return
	} else {
		*&r.Customers = (r.Customers)[1:]
		return
	}
}
