package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99}

	maps := Map(ints, func(a int) int {
		return a * 2
	})
	fmt.Println(maps)
}

func Map[T1, T2 any](s []T1, fn func(T1) T2) []T2 {
	newSlice := make([]T2, len(s))
	for idx, v := range s {
		newSlice[idx] = fn(v)
	}
	return newSlice
}
