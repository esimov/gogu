package main

import (
	"errors"
	"fmt"
)

func main() {
	ints := []int{2, 1, 4, 12, 8, 10, 22, 2, 10}

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

	fmt.Println("==================Without")
	fmt.Println(Without[int, int](ints, 2, 1, 12))

	fmt.Println("==================Difference")
	fmt.Println(Difference[int, int](ints, []int{2, 10, 4}))

	fmt.Println("==================Merge")
	fmt.Println(Merge(ints, []int{2, 10, 4}, []int{2, 23, 2}))

	fmt.Println("==================Flatten")
	sl1 := []any{[]any{1.0, 2.0, []any{3.0, []float64{4, 5, 6}}}, 7.0}
	fl1, _ := Flatten[float64](sl1)
	fmt.Println(fl1)

	fmt.Println("==================Union")
	sl2 := []any{[]any{1, 2, []any{3, []int{4, 5, 6}}}, 7, []int{1, 2}}
	fl2, _ := Union[int](sl2)
	fmt.Println(fl2)
}

func Map[T1, T2 any](s []T1, fn func(T1) T2) []T2 {
	newSlice := make([]T2, len(s))

	for idx, v := range s {
		newSlice[idx] = fn(v)
	}

	return newSlice
}

func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
}

func ForEachRight[T any](s []T, fn func(T)) {
	for i := len(s) - 1; i >= 0; i-- {
		fn(s[i])
	}
}

func Reduce[T1, T2 any](s []T1, fn func(T1, T2) T2, initVal T2) T2 {
	actual := initVal

	for _, v := range s {
		actual = fn(v, actual)
	}

	return actual
}

func Reverse[T any](s []T) []T {
	rev := make([]T, len(s))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = s[j], s[i]
	}

	return rev
}

// Unique returns slice's unique values.
func Unique[T comparable](s []T) []T {
	keys := make(map[T]bool)
	uni := []T{}
	for _, v := range s {
		if _, ok := keys[v]; !ok {
			keys[v] = true
			uni = append(uni, v)
		}
	}
	return uni
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
func Difference[T1 comparable, T2 any](s1, s2 []T1) []T1 {
	keys := make(map[T1]bool)
	uni := []T1{}
loop:
	for _, v := range s1 {
		for _, val := range s2 {
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
