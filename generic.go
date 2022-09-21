package gogu

import (
	"golang.org/x/exp/constraints"
)

type CompFn[T any] func(a, b T) bool

func Compare[T comparable](a, b T, comp CompFn[T]) int {
	if comp(a, b) {
		return 1
	} else if comp(b, a) {
		return -1
	}
	return 0
}

func Equal[T comparable](a, b T) bool {
	return a == b
}

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}
