package main

import "fmt"

func main() {
	ints := sumMap[string, int64](map[string]int64{
		"one": 1,
		"two": 2,
	})

	floats := sumMap[string, float64](map[string]float64{
		"one": 1.1,
		"two": 2.2,
	})

	fmt.Println(ints)
	fmt.Println(floats)
}

type MapVal interface {
	int64 | float64
}

func sumMap[K comparable, V MapVal](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
