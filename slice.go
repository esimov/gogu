package gogu

import (
	"errors"
)

// Map produces a new slice of values by mapping each value in the list through a transformation function.
func Map[T, R any](s []T, fn func(T) R) []R {
	result := make([]R, len(s))

	for idx, v := range s {
		result[idx] = fn(v)
	}
	return result
}

// ForEach iterates over the elements of a collection and invokes the callback fn function on each element.
func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
}

// ForEachRight is the same as ForEach, but starts the iteration from the last element.
func ForEachRight[T any](s []T, fn func(T)) {
	for i := len(s) - 1; i >= 0; i-- {
		fn(s[i])
	}
}

// Reduce reduces the collection to a value which is the accumulated result of running
// each element in the collection through the callback function yielding a single value.
func Reduce[T1, T2 any](s []T1, fn func(T1, T2) T2, initVal T2) T2 {
	actual := initVal

	for _, v := range s {
		actual = fn(v, actual)
	}

	return actual
}

// Reverse reverses the order of elements so that the first element becomes the last,
// the second element becomes the second to last, and so on.
func Reverse[T any](s []T) []T {
	rev := make([]T, len(s))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = s[j], s[i]
	}

	return rev
}

// Unique returns the collection unique values.
func Unique[T comparable](s []T) []T {
	keys := make(map[T]bool)
	result := []T{}

	for _, v := range s {
		if _, ok := keys[v]; !ok {
			keys[v] = true
			result = append(result, v)
		}
	}

	return result
}

// Every returns true if all of the elements of a slice satisfies the criteria of the callback function.
func Every[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Some returns true if some of the elements of a slice satisfies the criteria of the callback function.
func Some[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

// Contains returns true if the value is present in the slice.
func Contains[T comparable](s []T, value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

// Duplicate returns the duplicated values of a collection.
func Duplicate[T comparable](s []T) []T {
	keyCount := make(map[T]int)
	result := []T{}

	// Count how many times a value is showing up in the provided collection.
	for _, v := range s {
		if _, ok := keyCount[v]; !ok {
			keyCount[v] = 1
		} else {
			keyCount[v]++
		}
	}

	// Include only the values which count frequency is greater than 1 into the resulting slice.
	for k, v := range keyCount {
		if v > 1 {
			result = append(result, k)
		}
	}
	return result
}

// DuplicateWithIndex returns the duplicated values of a collection and their corresponding position as map.
func DuplicateWithIndex[T comparable](s []T) map[T]int {
	var count int
	kv := make(map[T][]int)
	result := make(map[T]int)

	// Count how many times a value is showing up in the provided collection.
	for idx, v := range s {
		if _, ok := kv[v]; !ok {
			// Create a slice with a dimension of 2, which first element contains the position (the index)
			// of the first found duplicate value and the second indicates the number of appearance.
			kv[v] = make([]int, 2)
			count = 1
			kv[v][0] = idx
			kv[v][1] = count
		} else {
			count++
			kv[v][1] = count
		}
	}

	// Include only the values which count frequency is greater than 1 into the resulting slice.
	for k, v := range kv {
		if v[1] > 1 {
			result[k] = v[0]
		}
	}
	return result
}

// Merge merges the first slice with the other slices defined as variadic parameter.
func Merge[T any](s []T, slices ...[]T) []T {
	merged := []T{}

	for i := 0; i < len(slices); i++ {
		merged = append(merged, slices[i]...)
	}
	merged = append(s, merged...)

	return merged
}

// Flatten flattens the slice all the way to the deepest nesting level.
func Flatten[T any](s any) ([]T, error) {
	return baseFlatten([]T{}, s)
}

func baseFlatten[T any](acc []T, s any) ([]T, error) {
	var err error

	switch v := (any)(s).(type) {
	case T:
		acc = append(acc, v)
	case []T:
		acc = append(acc, v...)
	case []any:
		for _, sv := range v {
			acc, err = baseFlatten(acc, sv)
			if err != nil {
				return nil, errors.New("flattening error")
			}
		}
	default:
		return nil, errors.New("flattening error")
	}

	return acc, nil
}

// Union computes the union of the passed-in slice and returns in order the list
// of unique items that are present in one or more of the slices.
func Union[T comparable](s any) ([]T, error) {
	var err error
	if flatten, err := baseFlatten([]T{}, s); err == nil {
		return Unique(flatten), nil
	}
	return nil, err
}

// Intersection computes the list of values that are the intersection of all the slices.
func Intersection[T comparable](s any) ([]T, error) {
	var err error

	flatten, err := baseFlatten([]T{}, s)
	if err != nil {
		return nil, err
	}

	return Duplicate(flatten), nil
}

// IntersectionBy is like Intersection, except that it accepts and callback function which is invoked on each element of the collection.
func IntersectionBy[T comparable](fn func(T) T, slices ...[]T) ([]T, error) {
	merged, result := []T{}, []T{}

	for _, s := range slices {
		merged = append(merged, s...)
	}

	dups := DuplicateWithIndex(Map(merged, fn))
	for _, v := range dups {
		result = append(result, merged[v])
	}

	return result, nil
}

// Without returns a copy of the slice with all the values defined in the variadic parameter removed.
func Without[T1 comparable, T2 any](s []T1, values ...T1) []T1 {
	keys := make(map[T1]bool)
	uni := []T1{}
loop:
	for _, v := range s {
		for _, val := range values {
			if v == val {
				continue loop
			}
		}
		if _, ok := keys[v]; !ok {
			keys[v] = true
			uni = append(uni, v)
		}
	}

	return uni
}

// Difference is similar to Without, but returns the values from
// the first slice that are not present in the second slice.
func Difference[T comparable](s1, s2 []T) []T {
	keys := make(map[T]bool)
	unique := []T{}
loop:
	for _, v := range s1 {
		for _, val := range s2 {
			if v == val {
				continue loop
			}
		}
		if _, ok := keys[v]; !ok {
			keys[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
