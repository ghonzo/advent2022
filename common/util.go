package common

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value
func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Atoi is just like the one in strconv, except we throw out the error
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Sgn returns 1 for a positive number, -1 for a negative number, and 0 for zero
func Sgn[T constraints.Signed | constraints.Float](a T) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return 1
	}
	return 0
}
