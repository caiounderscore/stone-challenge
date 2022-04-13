package grocery

import (
	"reflect"
	"testing"
)

func TestNewGrocery(t *testing.T) {
	type args struct {
		rs []*Register
		cs []*Customer
	}
	tests := []struct {
		name string
		args args
		want *Grocery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrocery(tt.args.rs, tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrocery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrocery_SortCustomers(t *testing.T) {
	tests := []struct {
		name string
		g    *Grocery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.SortCustomers()
		})
	}
}

func TestGrocery_ProcessItem(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		g    *Grocery
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.ProcessItem(tt.args.t)
		})
	}
}
