package main

import (
	"errors"
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

func main() {
	ints := []int{2, 1, 4, 12, 8, 10, 22, 2, 10, 2, 13, 10, 4, 13}

	fmt.Println("==================Map")
	maps := Map(ints, func(a int) int {
		return a * 2
	})
	fmt.Println(maps)

	fmt.Println("==================ForEach")
	ForEach(ints, func(v int) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()

	fmt.Println("==================ForEachRight")
	ForEachRight(ints, func(v int) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()

	fmt.Println("==================Reduce")
	reduce := Reduce(ints, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(reduce)

	fmt.Println("==================Reverse")
	fmt.Println(Reverse(ints))

	fmt.Println("==================Unique")
	fmt.Println(Unique[int](ints))

	fmt.Println("==================Duplicate")
	fmt.Println(Duplicate[int](ints))

	fmt.Println("==================Duplicate With Index")
	fmt.Println(DuplicateWithIndex[int](ints))

	fmt.Println("==================Duplicate Strings")
	strs := []string{"One", "Two", "Foo", "Bar", "Baz", "Foo", "Foo", "One"}
	fmt.Println(Duplicate(strs))

	fmt.Println("==================Merge")
	fmt.Println(Merge(ints, []int{2, 10, 4}, []int{2, 23, 2}))

	fmt.Println("==================Without")
	fmt.Println(Without[int, int](ints, 2, 1, 12))

	fmt.Println("==================Difference")
	fmt.Println(Difference[int](ints, []int{2, 10, 4}))

	sl1 := []any{[]any{1.0, 2.0, []any{3.0, []float64{4, 5, 6}}}, 7.0}

	fmt.Println("==================Flatten")
	fl, _ := Flatten[float64](sl1)
	fmt.Println(fl)

	sl2 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 10, 10}

	fmt.Println("==================Union")
	un, _ := Union[int](sl2)
	fmt.Println(un)

	str2 := []any{[]any{"One", "Two", []any{"Foo", []string{"Bar", "Baz", "Qux"}}}, "Foo", []string{"Foo", "Two"}, "Baz", "bar"}

	fmt.Println("==================Union Strings")
	sl3, _ := Union[string](str2)
	fmt.Println(sl3)

	fmt.Println("==================Intersection")
	sl4 := []any{[]int{1, 2, 3}, []int{101, 2, 1, 10}, []int{2, 1}}
	in, _ := Intersection[int](sl4)
	fmt.Println(in)

	fmt.Println("==================IntersectionBy")
	fl4, _ := IntersectionBy[float64](func(v float64) float64 {
		return math.Floor(v)
	}, []float64{2.1, 1.2, 5.09}, []float64{2.3, 2.2, 3.04, 3.1, 4.8, 4.1})
	fmt.Println(fl4)

	fmt.Println("==================Range")
	rn, err := Range[int](-10, -2, -320)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rn)
	}

	fmt.Println("==================RangeRight")
	rr, err := RangeRight[int](0, -1, -4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rr)
	}
}

// Map produces a new slice of values by mapping each value in list through a transformation function.
func Map[T, R any](s []T, fn func(T) R) []R {
	result := make([]R, len(s))

	for idx, v := range s {
		result[idx] = fn(v)
	}
	return result
}

// ForEach iterates over the elements of a collection and invokes fn for each element.
func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
}

// ForEachRight is the same as ForEach, but this starts the iteration from the last element.
func ForEachRight[T any](s []T, fn func(T)) {
	for i := len(s) - 1; i >= 0; i-- {
		fn(s[i])
	}
}

// Reduce reduces the collection to a value which is the accumulated result of running
// each element in collection through the iteratee function yielding a single value.
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
		for _, v := range slices[i] {
			merged = append(merged, v)
		}
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

// IntersectionBy is like Intersection, except that it accepts and iteratee function which is invoked on each element of the collection.
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
func Difference[T1 comparable](s1, s2 []T1) []T1 {
	keys := make(map[T1]bool)
	unique := []T1{}
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

// Range creates a slice of numbers (integers) progressing from start (if omitted defaults to 0) until the end.
// This method can accept 1, 2 or 3 parameters. Depending on the number of provided parameters, `start`, `step` and `end` are having the following meanings:
// [start=0]: The start of the range.
// [step=1]: The value to increment or decrement by.
// end: The end of the range.

// In case you'd like negative values, use a negative step.
// TODO make a thorough test.
func Range[T ~int](params ...T) ([]T, error) {
	var result []T

	if len(params) > 3 {
		return nil, errors.New("the method require maximum 3 paramenters.")
	}

	var start, step, end T

	switch len(params) {
	case 1:
		step = 1
		end = params[len(params)-1]
	case 2:
		start = params[0]
		step = 1
		end = params[len(params)-1]
	case 3:
		start = params[0]
		step = params[1]
		end = params[len(params)-1]

		if step == 0 {
			return nil, errors.New("step value should not be zero.")
		}
		if step < 0 && end > start {
			return nil, errors.New("the end value should be less than the start value in case you are using a negative increment.")
		}
	default:
		return nil, errors.New("the method require at least one paramenter, which should be the range dimension in this case.")
	}

	if end > 0 {
		for i := start; i < end; i += step {
			result = append(result, i)
		}
	} else {
		for i := start; end < i; i -= Abs(step) {
			result = append(result, i)
		}
	}

	return result, nil
}

// RangeRight is like Range, only that it populates the slice in descending order.
func RangeRight[T ~int](params ...T) ([]T, error) {
	ran, err := Range(params...)
	if err != nil {
		return nil, err
	}
	return Reverse(ran), nil
}

// Min returns the slowest value of the provided parameters.
func Min[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v < acc {
			acc = v
		}
	}
	return acc
}

// Max returns the biggest value of the provided parameters.
func Max[T constraints.Ordered](values ...T) T {
	var acc T = values[0]

	for _, v := range values {
		if v > acc {
			acc = v
		}
	}
	return acc
}

// Abs returns the absolut value of x.
func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
