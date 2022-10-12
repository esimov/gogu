package torx

import (
	"fmt"
	"sync"
	"time"

	"github.com/esimov/torx/cache"
	"golang.org/x/exp/constraints"
)

// Flip creates a function that invokes fn with arguments reversed.
func Flip[T any](fn func(args ...T) []T) func(args ...T) []T {
	return func(args ...T) []T {
		return Reverse(fn(args...))
	}
}

// Delay invokes the function with a predefined delay.
func Delay(delay time.Duration, fn func()) *time.Timer {
	t := time.AfterFunc(delay, fn)
	return t
}

// After creates a function wrapper that does nothing at first.
// From the nth call onwards, it starts actually calling the callback function.
// Useful for grouping responses, where you want to be sure that all the calls have finished, before proceeding.
func After[V constraints.Signed](n *V, fn func()) {
	if *n < 1 {
		fn()
	}
	*n-- // decrease the n as pointer receiver
}

// Before creates a function wrapper that memoizes its return value.
// From the nth call onwards, the memoized result of the last invocation is returned immediately
// instead of invoking function again. So the wrapper will invoke function at most n-1 times.
func Before[S ~string, T any, V constraints.Signed](n *V, c *cache.Cache[S, T], fn func() T) T {
	var memo *cache.Item[T]
	*n-- // decrease the n as pointer receiver
	if *n > 0 {
		return fn()
	}
	if *n == 0 {
		c.Set("func", fn(), cache.DefaultExpiration)
	}
	memo, _ = c.Get("func")

	return memo.Val()
}

var n = 2

// Once is like Before, but it's invoked only once.
// Repeated calls to the modified function will have no effect, returning the value from the cache.
func Once[S ~string, T any](c *cache.Cache[S, T], fn func() T) T {
	return Before(&n, c, fn)
}

// RType is a generic struct type used as method receiver on retry operations.
type RType[T any] struct {
	Input T
}

// Retry tries to invoke the callback function n times.
// It runs until the number of attempts is reached or the returned value of the callback function is nil.
func (v RType[T]) Retry(n int, fn func(T) error) (int, error) {
	var (
		err     error
		attempt int
	)

	if n < 0 {
		return attempt, fmt.Errorf("the number of attempts should be a positive number, got %v", n)
	}

	for attempt < n {
		if err = fn(v.Input); err == nil {
			return attempt, nil
		}
		attempt++
	}

	return attempt, err
}

// RetryWithDelay tries to invoke the callback function n times, but with a delay between each calls.
// It runs until the number of attempts is reached or the error return value of the callback function is nil.
func (v RType[T]) RetryWithDelay(n int, delay time.Duration, fn func(time.Duration, T) error) (time.Duration, int, error) {
	var (
		err     error
		attempt int
	)

	start := time.Now()
	for attempt < n {
		err = fn(time.Since(start), v.Input)
		if err == nil {
			return time.Since(start), attempt, nil
		}
		<-time.After(delay)
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
// until the time duration has elapsed since the last invocation passed in as a function argument.
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

// The throttle implementation is based on this package: https://github.com/boz/go-throttle.
type throttler struct {
	duration time.Duration
	cond     *sync.Cond
	last     time.Time
	waiting  bool
	trailing bool
	stop     bool
}

// NewThrottle creates a throttled function in order to limit the frequency rate at which the passed in function is invoked.
// The throttled function comes with a cancel method for canceling delayed function invocation.
// If the trailing parameter is true, the function is invoked right after the throttled code
// has been started, but at the trailing edge of the timeout.
// In this case the code will be executed one more time at the beginning of the next period.
//
// This function is useful for rate-limiting events that occur faster than you can keep up with.
func NewThrottle(wait time.Duration, trailing bool) *throttler {
	t := &throttler{
		cond: &sync.Cond{
			L: new(sync.Mutex),
		},
		duration: wait,
		trailing: trailing,
	}
	return t
}

// Call schedules the execution of the passed in function after the predefined delay.
func (t *throttler) Call() {
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
func (t *throttler) Next() bool {
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
func (t *throttler) Cancel() {
	t.cond.L.Lock()
	defer t.cond.L.Unlock()

	t.stop = true
	t.cond.Broadcast()
}
