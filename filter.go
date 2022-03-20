package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99}

	result := Filter(ints, func(a int) bool {
		return a > 10
	})
	fmt.Println(result)

	floats := []float64{12.2, 23.1, 10.01, 1, 643, 99}
	result2 := Filter(floats, func(a float64) bool {
		return a > 10
	})
	fmt.Println(result2)
}

func Filter[T any](s []T, fn func(T) bool) []T {
	rs := make([]T, 0)

	for _, v := range s {
		if fn(v) {
			rs = append(rs, v)
		}
	}

	return rs
}
