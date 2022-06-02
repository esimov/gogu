package main

import (
	"errors"
	"fmt"
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
		fmt.Println(v)
	})

	fmt.Println("==================ForEachRight")
	ForEachRight(ints, func(v int) {
		fmt.Println(v)
	})

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
	fl1, _ := Flatten[float64](sl1)
	fmt.Println(fl1)

	sl2 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}, 3, []int{4, 7}, 10, 10}

	fmt.Println("==================Union")
	fl2, _ := Union[int](sl2)
	fmt.Println(fl2)

	str2 := []any{[]any{"One", "Two", []any{"Foo", []string{"Bar", "Baz", "Qux"}}}, "Foo", []string{"Foo", "Two"}, "Baz", "bar"}

	fmt.Println("==================Union Strings")
	fls2, _ := Union[string](str2)
	fmt.Println(fls2)

	// fmt.Println("==================Intersection")
	// fl3, _ := Intersection[int](sl2)
	// fmt.Println(fl3)

	// fmt.Println("==================IntersectionBy")
	// fl4, _ := IntersectionBy[float64](func(v float64) float64 {
	// 	//fmt.Println(math.Floor(v))
	// 	return math.Floor(v)
	// }, []float64{2.1, 1.2, 4.2}, []float64{2.3, 2.2, 3.04})
	// fmt.Println(fl4)
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
	keys := make(map[T]int)
	result := []T{}

	// Count how many times is showing up a value in the provided collection.
	for _, v := range s {
		if _, ok := keys[v]; !ok {
			keys[v] = 0
		} else {
			keys[v]++
		}
	}

	// Include the values which count frequency is greater than 0 into the resulting slice.
	for k, v := range keys {
		if v > 0 {
			result = append(result, k)
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

func IntersectionBy[T comparable](fn func(T) T, slices ...[]T) ([]T, error) {
	merged, result := []T{}, []T{}

	for _, s := range slices {
		merged = append(merged, s...)
	}

	dup := Duplicate(Map(merged, fn))
	fmt.Println(dup)
	for k, _ := range dup {
		result = append(result, merged[k])
	}

	return result, nil
}

// func IntersectionBy[T comparable](s any, fn func(T) bool) ([]T, error) {
// 	res := []T{}
// 	flatten, err := baseFlatten([]T{}, s)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, v := range flatten {
// 		if fn(v) {
// 			res = append(res, v)
// 		}
// 	}
// 	return res, nil
// }

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
