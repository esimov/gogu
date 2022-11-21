package torx

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// FindIndex returns the index of the first found element.
func FindIndex[T any](s []T, fn func(T) bool) int {
	for k, v := range s {
		if fn(v) {
			return k
		}
	}
	return -1
}

// FindLastIndex is like FindIndex, only that returns the index of last found element.
func FindLastIndex[T any](s []T, fn func(T) bool) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[i]) {
			return i
		}
	}
	return -1
}

// FindAll is like FindIndex, but returns into a map all the values
// which stisfies the conditional logic of the callback function.
// The map key represents the position of the found value and the value is the item itself.
func FindAll[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for k, v := range s {
		if fn(v) {
			m[k] = v
		}
	}

	return m
}

// FindMin finds the minumum value of a slice.
func FindMin[T constraints.Ordered](s []T) T {
	var min T
	if len(s) > 0 {
		min = s[0]
	}
	for i := 0; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

// FindMinBy is like FindMin except that it accept a callback function
// and the conditional logic is applied over the resulted value.
// If there are more than one identical values resulted
// from the callback function the first one is used.
func FindMinBy[T constraints.Ordered](s []T, fn func(val T) T) T {
	var min T
	if len(s) > 0 {
		min = s[0]
	}

	for i := 0; i < len(s); i++ {
		if fn(s[i]) < fn(min) {
			min = s[i]
		}
	}
	return min
}

// FindMinByKey finds the minimum value from a map by using some existing key as a parameter.
func FindMinByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) (T, error) {
	var min T
	if _, ok := mapSlice[0][key]; !ok {
		return min, errors.New("key not found")
	}

	if len(mapSlice) > 0 {
		min = mapSlice[0][key]
	}

	for _, m := range mapSlice {
		mapped := FindByKey(m, func(k K) bool {
			return k == key
		})
		if _, ok := mapped[key]; ok {
			if mapped[key] < min {
				min = mapped[key]
			}
		}
	}

	return min, nil
}

// FindMax finds the maximum value of a slice.
func FindMax[T constraints.Ordered](s []T) T {
	var max T
	if len(s) > 0 {
		max = s[0]
	}
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// FindMaxBy is like FindMax except that it accept a callback function
// and the conditional logic is applied over the resulted value.
// If there are more than one identical values resulted
// from the callback function the first one is returned.
func FindMaxBy[T constraints.Ordered](s []T, fn func(val T) T) T {
	var max T
	if len(s) > 0 {
		max = s[0]
	}

	for i := 0; i < len(s); i++ {
		if fn(s[i]) > fn(max) {
			max = s[i]
		}
	}
	return max
}

// FindMaxByKey finds the maximum value from a map by using some existing key as a parameter.
func FindMaxByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) (T, error) {
	var max T
	if _, ok := mapSlice[0][key]; !ok {
		return max, errors.New("key not found")
	}

	if len(mapSlice) > 0 {
		max = mapSlice[0][key]
	}

	for _, m := range mapSlice {
		mapped := FindByKey(m, func(k K) bool {
			return k == key
		})
		if _, ok := mapped[key]; ok {
			if mapped[key] > max {
				max = mapped[key]
			}
		}
	}

	return max, nil
}

// Nth returns the nth element of the collection.
// In case of negative value the nth element is returned from the end of the collection.
// In case nth is out of bounds an error is returned.
func Nth[T any](slice []T, nth int) (T, error) {
	bounds := Bound[int]{0, len(slice)}

	if (nth > 0 && nth > bounds.Max-1) ||
		(nth < 0 && bounds.Max-Abs(nth) < 0) {

		var t T
		return t, fmt.Errorf("%d out of slice bounds %d", nth, bounds.Max)
	}

	if bounds.Enclose(nth) {
		if nth >= 0 {
			return slice[nth], nil
		}
	}
	return slice[len(slice)-Abs(nth)], nil
}

type Bound[T constraints.Signed] struct {
	Min, Max T
}

// Enclose checks if an element is inside the bounds.
func (b Bound[T]) Enclose(nth T) bool {
	if Abs(nth) >= b.Min && Abs(nth) <= b.Max {
		return true
	}

	return false
}
