package main

import "fmt"

func main() {
	ints := []int{12, 23, 1, 643, 99}
	reduce := Reduce(ints, func(a, b int) int {
		return a + b
	}, 0)
	fmt.Println(reduce)
}

func Reduce[T1, T2 any](s []T1, fn func(T1, T2) T2, initVal T2) T2 {
	actual := initVal
	for _, v := range s {
		actual = fn(v, actual)
	}
	return actual
}
