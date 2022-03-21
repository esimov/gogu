package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99, 2}
	reduce := Reduce(ints, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(reduce)

	ForEach(ints, func(v int) {
		fmt.Println(v)
	})

	res1 := Find(ints, func(v int) bool {
		return v == 12
	})

	res2 := FindLast(ints, func(v int) bool {
		return v == 12
	})

	fmt.Println("Find first:", res1)
	fmt.Println("Find last:", res2)

	fmt.Println(Reverse(ints))
	fmt.Println(IndexOf(ints, 12))
}

func Reduce[T1, T2 any](s []T1, fn func(T1, T2) T2, initVal T2) T2 {
	actual := initVal

	for _, v := range s {
		actual = fn(v, actual)
	}

	return actual
}

func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
}

func ForEachRight[T any](s []T, fn func(T)) {
	for i := len(s) - 1; i > 0; i-- {
		fn(s[i])
	}
}

func Find[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, 1)

	for k, v := range s {
		if fn(v) {
			m[k] = v
			return m
		}
	}

	return nil
}

func FindLast[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, 1)

	for i := len(s) - 1; i > 0; i-- {
		if fn(s[i]) {
			m[i] = s[i]
			return m
		}
	}

	return nil
}

func Reverse[T any](s []T) []T {
	var rs = make([]T, len(s))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = s[j], s[i]
	}

	return rs
}

func IndexOf[T comparable](s []T, value T) int {
	for k, v := range s {
		if v == value {
			return k
		}
	}
	return -1
}
