package grocery

import (
	"reflect"
	"testing"
)

func TestNewRegister(t *testing.T) {
	type args struct {
		id int
		t  bool
	}
	tests := []struct {
		name string
		args args
		want *Register
	}{
		{
			name: "Test New Register",
			args: args{
				id: 1,
				t:  false,
			},
			want: &Register{
				Id:        1,
				Training:  false,
				Customers: make([]Customer, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegister(tt.args.id, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildRegisters(t *testing.T) {

	type args struct {
		r []int
	}
	tests := []struct {
		name string
		args args
		want []*Register
	}{
		{
			name: "Build Registers",
			args: args{
				r: []int{2},
			},
			want: []*Register{
				{
					Id:        1,
					Customers: make([]Customer, 0),
					Training:  false,
				}, {
					Id:        2,
					Customers: make([]Customer, 0),
					Training:  true,
				},
			},
		},
		{
			name: "Build Registers - Case 2",
			args: args{
				r: []int{1},
			},
			want: []*Register{
				{
					Id:        1,
					Customers: make([]Customer, 0),
					Training:  true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRegisters(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRegisters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSmallRegister(t *testing.T) {
	type args struct {
		rs []*Register
	}
	tests := []struct {
		name string
		args args
		want *Register
	}{
		{
			name: "Get Small Register",
			args: args{
				rs: []*Register{
					{
						Id:       1,
						Training: false,
						Customers: []Customer{
							{
								Id:    1,
								Items: 5,
							},
							{
								Id:    2,
								Items: 4,
							},
						},
					},
					{
						Id:       1,
						Training: false,
						Customers: []Customer{
							{
								Id:    1,
								Items: 5,
							},
						},
					},
				},
			},
			want: &Register{
				Id:       1,
				Training: false,
				Customers: []Customer{
					{
						Id:    1,
						Items: 5,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallRegister(tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSmallRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSmallRegisterItemRecord(t *testing.T) {
	type args struct {
		rs []*Register
	}
	tests := []struct {
		name string
		args args
		want *Register
	}{
		{
			name: "Get Small Register Item Record",
			args: args{
				rs: []*Register{
					{
						Id:       1,
						Training: false,
						Customers: []Customer{
							{
								Id:    1,
								Items: 5,
							},
							{
								Id:    2,
								Items: 4,
							},
						},
					},
					{
						Id:       1,
						Training: false,
						Customers: []Customer{
							{
								Id:    1,
								Items: 5,
							},
						},
					},
				},
			},
			want: &Register{
				Id:       1,
				Training: false,
				Customers: []Customer{
					{
						Id:    1,
						Items: 5,
					},
					{
						Id:    2,
						Items: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallRegisterItemRecord(tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSmallRegisterItemRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckRegistersIfEmpty(t *testing.T) {
	type args struct {
		rs []*Register
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Check Registers if Empty",
			args: args{
				rs: []*Register{
					{
						Id:       1,
						Training: false,
						Customers: []Customer{
							{
								Id:    1,
								Items: 5,
							},
							{
								Id:    1,
								Items: 4,
							},
						},
					},
					{
						Id:       1,
						Training: false,
						Customers: []Customer{
							{
								Id:    2,
								Items: 5,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckRegistersIfEmpty(tt.args.rs); got != tt.want {
				t.Errorf("CheckRegistersIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister_Push(t *testing.T) {
	type args struct {
		c Customer
	}
	tests := []struct {
		name string
		r    *Register
		args args
	}{
		{
			name: "Test Register Push",
			r: &Register{
				Id:        1,
				Customers: make([]Customer, 0),
				Training:  false,
			},
			args: args{
				c: Customer{
					Id: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Push(tt.args.c)
			if tt.r.Customers[0] != tt.args.c {
				t.Errorf("Register.Size() = %v, want %v", tt.r.Customers[0], tt.args.c)

			}
		})
	}
}

func TestRegister_Size(t *testing.T) {
	tests := []struct {
		name string
		r    *Register
		want int
	}{
		{
			name: "Test Register Size",
			r: &Register{
				Id:       1,
				Training: false,
				Customers: []Customer{
					{
						Id:    1,
						Items: 5,
					},
					{
						Id:    2,
						Items: 4,
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Size(); got != tt.want {
				t.Errorf("Register.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister_Pop(t *testing.T) {
	tests := []struct {
		name string
		r    *Register
	}{

		{
			name: "Test Register Size",
			r: &Register{
				Id:       1,
				Training: false,
				Customers: []Customer{
					{
						Id:    1,
						Items: 5,
					},
					{
						Id:    2,
						Items: 4,
					},
				},
			},
		},
	}

	c := Customer{
		Id:    2,
		Items: 4,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.Pop()
			if tt.r.Customers[0] != c {
				t.Errorf("Register.Size() = %v, want %v", tt.r.Customers[0], c)

			}
		})
	}
}
