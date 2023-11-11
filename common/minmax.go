package common

import "golang.org/x/exp/constraints"

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
