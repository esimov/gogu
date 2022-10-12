package torx

import (
	"golang.org/x/exp/constraints"
)

// CompFn is a generic function type for comparing two values.
type CompFn[T any] func(a, b T) bool

// Compare compares two values using as comparator the the callback function argument.
func Compare[T comparable](a, b T, comp CompFn[T]) int {
	if comp(a, b) {
		return 1
	} else if comp(b, a) {
		return -1
	}
	return 0
}

// Equal checks if two values are equal.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Less checks if the first value is less than the second.
func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}
