package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99, 12, 2, 1}

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
	fmt.Println(Unique[int, int](ints))
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
func Unique[T1 comparable, T2 any](s []T1) []T1 {
	keys := make(map[T1]bool)
	uni := []T1{}
	for _, v := range s {
		if _, ok := keys[v]; !ok {
			keys[v] = true
			uni = append(uni, v)
		}
	}
	return uni
}
