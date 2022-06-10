package main

import (
	"fmt"

	"github.com/esimov/gogu"
)

func main() {
	fmt.Println("==================Clamp")
	fmt.Println(gogu.Clamp(10, -5, 5))

	fmt.Println("==================SumBy")
	fmt.Println(gogu.SumBy[int]([]int{1, 2, 3}, func(val int) int {
		return val * val
	}))
}
