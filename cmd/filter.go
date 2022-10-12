package main

import (
	"fmt"

	"github.com/esimov/torx"
)

func main() {
	ints := []int{12, 23, 1, 643, 99}

	fmt.Println("==================Filter")
	result := torx.Filter(ints, func(a int) bool {
		return a > 10
	})
	fmt.Println(result)

	floats := []float64{12.2, 23.1, 10.01, 1, 643, 99}
	result2 := torx.Filter(floats, func(a float64) bool {
		return a > 10
	})
	fmt.Println(result2)

	fmt.Println("==================Reject")
	fmt.Println(torx.Reject([]int{1, 2, 3, 4, 5, 6, 10, 20, 30, 40, 50}, func(num int) bool {
		return num%2 == 0
	}))

	mp := map[int]string{1: "John", 2: "Doe", 3: "Fred"}

	fmt.Println("==================FilterMap")
	res2 := torx.FilterMap[int, string](mp, func(v string) bool {
		return v == "John"
	})
	fmt.Println(res2)

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

	fmt.Println("==================Filter2DMapCollection")
	res3 := torx.Filter2DMapCollection[string, int](usersSlice, func(v map[string]int) bool {
		return v["age"] > 20 && v["ranking"] < 5
	})
	fmt.Println(res3)

	fmt.Println("==================FilterMapCollection")
	users := []map[string]int{
		{"bernie": 22},
		{"robert": 30},
	}
	res4 := torx.FilterMapCollection[string, int](users, func(val int) bool {
		return val > 22
	})
	fmt.Println(res4)
}
