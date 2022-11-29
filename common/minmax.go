package common

import "golang.org/x/exp/constraints"

// Max returns the greater of two values
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min returns the lesser of two values
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MaxElement returns the maximum value in the slice, , or "zero" if the slice is empty
func MaxElement[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}
	return m
}

// MinElement returns the minimum value in the slice, or "zero" if the slice is empty
func MinElement[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m > v {
			m = v
		}
	}
	return m
}

// MaxMin holds onto a maximum and minumum value of an arbitrary number of values.
// Create with new() and call the Accept method.
type MaxMin[T constraints.Ordered] struct {
	Max, Min    T
	initialized bool // will be false until we get at least one value
}

// Accept a new value and update the max and min according
func (mm *MaxMin[T]) Accept(v T) *MaxMin[T] {
	if v > mm.Max || !mm.initialized {
		mm.Max = v
	}
	if v < mm.Min || !mm.initialized {
		mm.Min = v
	}
	mm.initialized = true
	return mm
}
