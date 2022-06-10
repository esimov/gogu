package gogu

import (
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

// FindLastIndex is like FindIndex, only that iterates the slice in reverse order.
func FindLastIndex[T any](s []T, fn func(T) bool) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[j]) {
			return i
		}
	}
	return -1
}

// FindAll is like FindIndex, but put into a map all the values which stisfy the conditional logic of the callback function.
// The map key represents the position of the found value and the map value is the number itself.
func FindAll[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for k, v := range s {
		if fn(v) {
			m[k] = v
		}
	}

	return m
}

// FindAllFromLast is like FindAll, but run the slice iteration in backward order.
func FindAllFromLast[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[i]) {
			m[i] = s[j]
		}
	}

	return m
}

// IndexOf returns the index at which value can be found in the slice, or -1 if value is not present in the slice.
func IndexOf[T comparable](s []T, val T) int {
	for k, v := range s {
		if v == val {
			return k
		}
	}

	return -1
}

// LastIndexOf returns the index of the last occurrence of value.
func LastIndexOf[T comparable](s []T, val T) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if s[i] == val {
			return j
		}
	}
	return -1
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
func FindMinBy[T constraints.Ordered](s []T, fn func(val T) T) T {
	var min T
	if len(s) > 0 {
		min = fn(s[0])
	}

	for i := 0; i < len(s); i++ {
		if s[i] < fn(min) {
			min = s[i]
		}
	}
	return min
}

// FindMinByKey finds the minimum value from a map by using some existing key as a parameter.
func FindMinByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) T {
	var min T
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

	return min
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
func FindMaxBy[T constraints.Ordered](s []T, fn func(val T) T) T {
	var min T
	if len(s) > 0 {
		min = fn(s[0])
	}

	for i := 0; i < len(s); i++ {
		if s[i] < fn(min) {
			min = s[i]
		}
	}
	return min
}

// FindMaxByKey finds the maximum value from a map by using some existing key as a parameter.
func FindMaxByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) T {
	var max T
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

	return max
}
