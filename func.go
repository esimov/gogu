package gogu

import (
	"time"

	"golang.org/x/exp/constraints"
)

func Delay(delay time.Duration, fn func()) *time.Timer {
	t := time.AfterFunc(delay, fn)
	return t
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
// instead of invoking function again. So the wrapper will invoke function at most n-1 times.
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

// Once is like Before, but it's invoked only once.
func Once[T any, V constraints.Signed](n V, fn func() T) T {
	var memo *T = new(T)
	if n > 0 {
		*memo = fn()
	}
	if n <= 1 {
		fn()
		memo = new(T)
	}
	n-- // decrease the n through the pointer receiver
	return *memo
}

type RetryTyp[T any] struct {
	In T
}

// Retry tries to invoke the callback function n times.
// It runs until the number of attempts is reached or the callback function error return value is nil.
// In case n is less then 0, it runs until a successful response or no error is returned.
func (v RetryTyp[T]) Retry(n int, fn func(T) error) (int, error) {
	var (
		err     error
		attempt int
	)

	for attempt < n || n < 0 {
		if err = fn(v.In); err == nil {
			return attempt, nil
		}
		attempt++
	}

	return n, err
}

//func (v RetryTyp[T]) RetryWithDelay(n int, fn func(time.Duration))

// TODO remove
type fn[T any] func() T

func getf[T any](f fn[T]) T {
	return f()
}

type FuncType[T any] string

func (s FuncType[T]) get() string {
	return string(s)
}
