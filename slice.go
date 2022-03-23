package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99, 2}

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

	fmt.Println("==================IndexOf")
	fmt.Println(IndexOf(ints, 12))

	fmt.Println("==================LastIndexOf")
	fmt.Println(LastIndexOf(ints, 99))
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
	var rs = make([]T, len(s))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = s[j], s[i]
	}

	return rs
}

func IndexOf[T comparable](s []T, val T) int {
	for k, v := range s {
		if v == val {
			return k
		}
	}

	return -1
}

func LastIndexOf[T comparable](s []T, val T) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if s[i] == val {
			return j
		}
	}
	return -1
}
