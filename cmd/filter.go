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

	mp := map[int]string{1: "John", 2: "Doe", 3: "Fred"}

	fmt.Println("==================FilterMap")
	res2 := gogu.FilterMap[int, string](mp, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res2)

	usersMap := map[string]map[string]int{
		"bernie": {
			"age":     30,
			"ranking": 1,
		},

		"robert": {
			"age":     20,
			"ranking": 5,
		},
	}

	fmt.Println("==================Filter2DMap")
	res3 := gogu.Filter2DMap[string, int](usersMap, func(v int) bool {
		return v > 20
	})
	fmt.Println(res3)

	usersSlice := []map[string]map[string]int{
		{
			"bernie": {
				"age":     30,
				"ranking": 1,
			},
		},
		{
			"robert": {
				"age":     20,
				"ranking": 5,
			},
		},
	}

	fmt.Println("==================Filter2DMapSlice")
	res4 := gogu.Filter2DMapSlice[string, int](usersSlice, func(v map[string]int) bool {
		return v["age"] > 20 && v["ranking"] < 5
	})
	fmt.Println(res4)
}
