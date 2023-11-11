package common

import (
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
