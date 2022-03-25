package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99, 2, 2}

	fmt.Println("==================FindAll")
	res1 := FindAll(ints, func(v int) bool {
		return v == 2
	})
	fmt.Println(res1)

	fmt.Println("==================FindAllFromLast")
	res2 := FindAllFromLast(ints, func(v int) bool {
		return v == 12
	})
	fmt.Println(res2)

	fmt.Println("==================FindIndex")
	res3 := FindIndex(ints, func(v int) bool {
		return v == 23
	})
	fmt.Println(res3)

	fmt.Println("==================FindLastIndex")
	res4 := FindLastIndex(ints, func(v int) bool {
		return v == 23
	})
	fmt.Println(res4)

	fmt.Println("==================IndexOf")
	fmt.Println(IndexOf(ints, 12))

	fmt.Println("==================LastIndexOf")
	fmt.Println(LastIndexOf(ints, 99))
}

func FindAll[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for k, v := range s {
		if fn(v) {
			m[k] = v
		}
	}

	return m
}

func FindAllFromLast[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[i]) {
			m[i] = s[j]
		}
	}

	return m
}

func FindIndex[T any](s []T, fn func(T) bool) int {
	for k, v := range s {
		if fn(v) {
			return k
		}
	}
	return -1
}

func FindLastIndex[T any](s []T, fn func(T) bool) int {
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		if fn(s[j]) {
			return i
		}
	}
	return -1
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
