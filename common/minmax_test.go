package common

import (
	"reflect"
	"testing"
)

func TestMaxMin_Accept(t *testing.T) {
	mm := new(MaxMin[int])
	if mm.Max != 0 {
		t.Error("MaxMin.Max not initialized to 0")
	}
	if mm.Min != 0 {
		t.Error("MaxMin.Min not initialized to 0")
	}
	mm.Accept(6)
	if mm.Max != 6 {
		t.Errorf("MaxMin.Accept expected Max = %v, got %v", 6, mm.Max)
	}
	if mm.Min != 6 {
		t.Errorf("MaxMin.Accept expected Min = %v, got %v", 6, mm.Min)
	}
	mm.Accept(-8)
	if mm.Max != 6 {
		t.Errorf("MaxMin.Accept expected Max = %v, got %v", 6, mm.Max)
	}
	if mm.Min != -8 {
		t.Errorf("MaxMin.Accept expected Min = %v, got %v", -8, mm.Min)
	}
	if mm.Accept(100) != mm {
		t.Error("MaxMin.Accept didn't return the same instance")
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first, both positive", args{7, 3}, 7},
		{"first, both negative", args{-7, -13}, -7},
		{"second, both postive", args{79, 213}, 213},
		{"second, both negative", args{-79, -13}, -13},
		{"equal", args{-2, -2}, -2},
		{"zeros", args{0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first, both positive", args{7, 3}, 3},
		{"first, both negative", args{-7, -13}, -13},
		{"second, both postive", args{79, 213}, 79},
		{"second, both negative", args{-79, -13}, -79},
		{"equal", args{-2, -2}, -2},
		{"zeros", args{0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxElement(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}}, 0},
		{"one element zero", args{[]int{0}}, 0},
		{"one element negative", args{[]int{-299}}, -299},
		{"one element positive", args{[]int{279}}, 279},
		{"ascending", args{[]int{-3, 0, 5, 5, 9, 14}}, 14},
		{"descending", args{[]int{3, 0, -5, -5, -9, -14}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxElement(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinElement(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{[]int{}}, 0},
		{"one element zero", args{[]int{0}}, 0},
		{"one element negative", args{[]int{-299}}, -299},
		{"one element positive", args{[]int{279}}, 279},
		{"ascending", args{[]int{-3, 0, 5, 5, 9, 14}}, -3},
		{"descending", args{[]int{3, 0, -5, -5, -9, -14}}, -14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinElement(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
