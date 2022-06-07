package main

import (
	"fmt"

	"github.com/esimov/gogu"
)

func main() {
	ints := []int{12, 23, 1, 643, 99}

	result := gogu.Filter(ints, func(a int) bool {
		return a > 10
	})
	fmt.Println(result)

	floats := []float64{12.2, 23.1, 10.01, 1, 643, 99}
	result2 := gogu.Filter(floats, func(a float64) bool {
		return a > 10
	})
	fmt.Println(result2)
}
