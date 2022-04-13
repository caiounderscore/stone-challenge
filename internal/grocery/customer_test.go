package grocery

import (
	"testing"
)

// func TestCustomer_SelectRegister(t *testing.T) {
// 	type args struct {
// 		rs []*Register
// 		t  int
// 	}
// 	tests := []struct {
// 		name string
// 		c    *Customer
// 		args args
// 	}{
// 		{
// 			name: "Test Select Register - Senary 1",
// 			c: &Customer{
// 				Id:    1,
// 				Type:  "A",
// 				Items: 2,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.c.SelectRegister(tt.args.rs, tt.args.t)
// 		})
// 	}
// }

func TestBuildCustomers(t *testing.T) {
	type args struct {
		rawCustomers [][]string
	}
	tests := []struct {
		name string
		args args
		want []*Customer
	}{
		{
			name: "Test Build Customers",
			args: args{
				rawCustomers: [][]string{{"A", "1", "3"}, {"A", "1", "5"}, {"B", "3", "1"}},
			},
			want: []*Customer{
				{
					Id:          1,
					Type:        "A",
					Items:       3,
					ArrivalTime: 1,
				},
				{
					Id:          2,
					Type:        "A",
					Items:       5,
					ArrivalTime: 1,
				},
				{
					Id:          3,
					Type:        "B",
					Items:       1,
					ArrivalTime: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := BuildCustomers(tt.args.rawCustomers)
			if len(tt.want) != len(c) {
				t.Errorf("NewCustomer() = %v, want %v", len(c), len(tt.want))
			}

			for i := 0; i < len(tt.want); i++ {
				if *tt.want[i] != *c[i] {
					t.Errorf("NewCustomer() = %v, want %v", *c[i], *tt.want[i])
				}
			}

		})
	}
}
