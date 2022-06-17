package gogu

import (
	"time"

	"golang.org/x/exp/constraints"
)

func Delay(delay time.Duration, fn func()) {
	time.AfterFunc(delay, fn)
}

// After creates a function wrapper that does nothing at first. From the nth call onwards, it starts actually calling the callback function.
// Useful for grouping responses, where you want to be sure that all the calls have finished, before proceeding.
func After[T any, V constraints.Signed](n *V, fn func()) {
	if *n < 1 {
		fn()
	}
	*n-- // decrease the n through the pointer receiver
}

// Before creates a function wrapper that memoizes its return value.
// From the nth call onwards, the memoized result of the last invocation is returned immediately
// instead of invoking function again. So the wrapper will invoke function at most count - 1 times.
func Before[T any, V constraints.Signed](n *V, fn func() T) T {
	var memo *T = new(T)
	if *n > 0 {
		*memo = fn()
	}
	if *n <= 1 {
		fn()
		memo = new(T)
	}
	*n-- // decrease the n through the pointer receiver
	return *memo
}
