package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99, 12}
	reduce := Reduce(ints, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(reduce)

	ForEach[int, int](ints, func(v int) {
		fmt.Println(v)
	})

	res := Find(ints, func(v int) bool {
		return v == 12
	})
	fmt.Println("Find:", res)
}

func Reduce[T1, T2 any](s []T1, fn func(T1, T2) T2, initVal T2) T2 {
	actual := initVal

	for _, v := range s {
		actual = fn(v, actual)
	}

	return actual
}

func ForEach[T1, T2 any](s []T1, fn func(T1)) {
	for _, v := range s {
		fn(v)
	}
}

func Find[T any](s []T, fn func(T) bool) map[int]any {
	m := make(map[int]any, 1)
	for k, v := range s {
		if fn(v) {
			m[k] = v
			return m
		}
	}
	return nil
}
