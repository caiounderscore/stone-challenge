package grocery

import (
	"reflect"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	tests := []struct {
		name string
		want *Customer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_SelectRegister(t *testing.T) {
	type args struct {
		rs []*Register
		t  int
	}
	tests := []struct {
		name string
		c    *Customer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SelectRegister(tt.args.rs, tt.args.t)
		})
	}
}

func TestBuildCustomers(t *testing.T) {
	type args struct {
		rawCustomers [][]string
	}
	tests := []struct {
		name string
		args args
		want []*Customer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCustomers(tt.args.rawCustomers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}
