package gogu

import (
	"sync"
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

// RType is a generic struct type used as method receiver on retry operations.
type RType[T any] struct {
	Input T
}

// Retry tries to invoke the callback function n times.
// It runs until the number of attempts is reached or the callback function error return value is nil.
// In case n is less then 0, it runs until a successful response or no error is returned.
func (v RType[T]) Retry(n int, fn func(T) error) (int, error) {
	var (
		err     error
		attempt int
	)

	for attempt < n || n < 0 {
		if err = fn(v.Input); err == nil {
			return attempt, nil
		}
		attempt++
	}

	return attempt, err
}

// RetryWithDelay tries to invoke the callback function n times, but with a delay between each calls.
// It runs until the number of attempts is reached or the callback function error return value is nil.
// In case n is less then 0, it runs until a successful response or no error is returned.
func (v RType[T]) RetryWithDelay(n int, delay time.Duration, fn func(time.Duration, T) error) (time.Duration, int, error) {
	var (
		err     error
		attempt int
	)

	start := time.Now()
	for attempt < n || n < 0 {
		err = fn(time.Since(start), v.Input)
		if err == nil {
			return time.Since(start), attempt, nil
		}
		time.Sleep(delay)
		attempt++
	}

	return time.Since(start), attempt, err
}

type debouncer struct {
	duration time.Duration
	timer    *time.Timer
	mu       sync.Mutex
}

// NewDebounce creates a new debounced version of the invoked function which will postpone the execution
// until the time duration passed in as a function argument has elapsed since the last invocation.
//
// It returns a callback function which will be invoked after the predefined delay and
// also a cancel function which should be invoked to cancel a scheduled debounce.
func NewDebounce(wait time.Duration) (func(f func()), func()) {
	d := &debouncer{duration: wait}
	return func(f func()) {
		d.add(f)
	}, d.cancel
}

// add method schedules the execution of the passed in function after a predefined delay.
func (d *debouncer) add(f func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.duration, f)
}

// cancel the execution of a scheduled debounce function.
func (d *debouncer) cancel() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
		d.timer = nil
	}
}

type throttler struct {
	duration time.Duration
	cond     *sync.Cond
	last     time.Time
	waiting  bool
	trailing bool
	stop     bool
}

// NewThrottle creates a throttled function, useful to limit the frequency rate at which the passed in function is called.
// The throttled function comes with a cancel method for canceling delayed function invocation.
// If the trailing parameter is true, the function is invoked right after the throttled code
// has been started, but at the trailing edge of the timeout.
// In this case the code will be executed one more time at the beginning of the next period.
func NewThrottle(wait time.Duration, trailing bool) (func(f func()), func()) {
	t := &throttler{
		cond: &sync.Cond{
			L: new(sync.Mutex),
		},
		duration: wait,
		trailing: trailing,
	}

	return func(f func()) {
		go func() {
			for t.next() {
				f()
			}
		}()
		t.add(f)
	}, t.cancel
}

// add method schedules the execution of the passed in function after a predefined delay.
func (t *throttler) add(fn func()) {
	t.cond.L.Lock()
	defer t.cond.L.Unlock()

	if !t.waiting && !t.stop {
		delta := time.Since(t.last)
		if delta > t.duration {
			t.waiting = true
			t.cond.Broadcast()
		} else if t.trailing {
			t.waiting = true
			time.AfterFunc(t.duration-delta, t.cond.Broadcast)
		}
	}
}

// next returns true at most once per time period. It runs until the throttled function is not canceled.
func (t *throttler) next() bool {
	t.cond.L.Lock()
	defer t.cond.L.Unlock()

	for !t.waiting && !t.stop {
		t.cond.Wait()
	}

	if !t.stop {
		t.waiting = false
		t.last = time.Now()
	}

	return !t.stop
}

// cancel the execution of a scheduled throttle function.
func (t *throttler) cancel() {
	t.cond.L.Lock()
	defer t.cond.L.Unlock()

	t.stop = true
	t.cond.Broadcast()
}
