package main

import (
	"fmt"
	"math"

	"github.com/esimov/gogu"
)

func main() {
	ints := []int{12, 23, 1, 643, 99, 2, 2}

	fmt.Println("==================FindIndex")
	res3 := gogu.FindIndex(ints, func(v int) bool {
		return v == 2
	})
	fmt.Println(res3)

	fmt.Println("==================FindLastIndex")
	res4 := gogu.FindLastIndex(ints, func(v int) bool {
		return v == 2
	})
	fmt.Println(res4)

	fmt.Println("==================FindAll")
	res1 := gogu.FindAll(ints, func(v int) bool {
		return v == 2
	})
	fmt.Println(res1)

	fmt.Println("==================IndexOf")
	fmt.Println(gogu.IndexOf(ints, 12))

	fmt.Println("==================LastIndexOf")
	fmt.Println(gogu.LastIndexOf(ints, 99))

	fmt.Println("==================FindMin")
	fmt.Println(gogu.FindMin([]int{-1, 10, 2, 4, 99}))

	fmt.Println("==================FindMinBy")
	fmt.Println(gogu.FindMinBy([]float64{1.1, 1.4, 2.5, 4.8, 10.5}, func(val float64) float64 {
		return math.Ceil(val)
	}))

	fmt.Println("==================FindMinByKey")
	mp := []map[string]int{
		{"n": 1},
		{"n": 2},
		{"n": 3},
		{"n": 10},
		{"n": -7},
	}
	fmt.Println(gogu.FindMinByKey(mp, "n"))

	fmt.Println("==================FindMaxByKey")
	fmt.Println(gogu.FindMaxByKey(mp, "n"))

	fmt.Println("==================FindMax")
	fmt.Println(gogu.FindMax([]int{-1, 10, 2, 4, 99}))

	fmt.Println("==================Nth")
	fmt.Println(gogu.Nth[int]([]int{1, 2, 3, 4}, 1))
}
