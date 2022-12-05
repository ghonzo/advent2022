package common

import (
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{60}, 60},
		{"zero", args{0}, 0},
		{"neg", args{-999}, 999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{"225"}, 225},
		{"zero", args{"0"}, 0},
		{"neg", args{"-10"}, -10},
		{"empty", args{""}, 0},
		{"invalid", args{"pickle"}, 0},
		{"spaces", args{" 33 "}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atoi(tt.args.s); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSgn(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{60}, 1},
		{"zero", args{0}, 0},
		{"neg", args{-999}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sgn(tt.args.a); got != tt.want {
				t.Errorf("Sgn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"one char", args{"A"}, "A"},
		{"two char", args{"AB"}, "BA"},
		{"three char", args{"ABC"}, "CBA"},
		{"four char", args{"ABCD"}, "DCBA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
